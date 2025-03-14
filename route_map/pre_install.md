BEGIN
    // Vytisknutí ASCII banneru
    VYPIŠ "banner s logem a nápisem"

    // -----------------------------
    // Načtení globálních funkcí a proměnných
    // -----------------------------
    SET scrDir = adresář, kde se nachází aktuální skript
    SOURCE (scrDir + "/global_fn.sh")
      // global_fn.sh obsahuje pomocné funkce (např. print_log, prompt_timer, nvidia_detect, atd.)

    // -----------------------------
    // Vyhodnocení předaných parametrů (přepínačů)
    // -----------------------------
    INICIALIZUJ příznaky:
      flg_Install, flg_Restore, flg_Service, flg_DryRun, flg_Shell, flg_Nvidia, flg_ThemeInstall
    PROČTI příkazové řádky (getopts)
      // Přiřaď hodnoty příslušným příznakům (např. -i pro instalaci, -r pro obnovení konfigurace, apod.)
    KONEC cyklu

    // Nastavení proměnných pro export a kontrola výchozích voleb
    EXPORT proměnné (např. HYDE_LOG, flg_DryRun, atd.)
    IF žádný argument nebyl předán THEN
         Nastav výchozí: flg_Install, flg_Restore, flg_Service = 1
    END IF

    // -----------------------------
    // PŘEDINSTALAČNÍ KROK
    // -----------------------------
    IF (flg_Install == 1 AND flg_Restore == 1) THEN
         VYPIŠ "pre-install banner"
         SPUSŤ (scrDir + "/install_pre.sh")
           // install_pre.sh obsahuje předinstalační rutiny (např. příprava systému)
    END IF

    // -----------------------------
    // INSTALACE BALÍČKŮ A KONFIGURACE
    // -----------------------------
    IF (flg_Install == 1) THEN
         VYPIŠ "instalační banner"

         // Příprava seznamu balíčků
         COPY (scrDir + "/pkg_core.lst") -> (scrDir + "/install_pkg.lst")
         SET trap pro uložení seznamu do logu při ukončení
         IF byl předán vlastní seznam balíčků THEN
              APPEND obsah vlastního seznamu do install_pkg.lst
         END IF
         APPEND oddělovač "#user packages" do install_pkg.lst

         // Detekce Nvidia a přidání ovladačů, pokud jsou povoleny
         IF (nvidia_detect zjistí Nvidia GPU) THEN
              IF (flg_Nvidia == 1) THEN
                   PROČTI soubory s názvy balíčků z /usr/lib/modules/*/pkgbase
                   APPEND každý (kernel-headers) do install_pkg.lst
                   APPEND výsledky (nvidia_detect --drivers) do install_pkg.lst
              ELSE
                   VYPIŠ, že Nvidia akce jsou přeskočeny
              END IF
         END IF
         VOLAT nvidia_detect --verbose

         // Získání uživatelských preferencí (AUR helper, shell)
         IF (AUR helper není nastaven) THEN
              VYPIŠ seznam dostupných AUR helperů
              VYZVÁNÍ k volbě s timeoutem
              NAZEV vybraného AUR helperu se exportuje do proměnné (např. "yay-bin")
         END IF

         IF (shell preference není nastavena) THEN
              VYPIŠ seznam podporovaných shellů (např. zsh, fish)
              VYZVÁNÍ k volbě s timeoutem
              EXPORT vybraného shellu do proměnné a APPEND do install_pkg.lst
         END IF

         // Kontrola, že install_pkg.lst obsahuje "#user packages"
         IF (install_pkg.lst neobsahuje "#user packages") THEN
              VYPIŠ chybu a ukonči skript
         END IF

         // Spuštění instalace balíčků, pokud nejde o testovací režim
         IF (flg_DryRun != 1) THEN
              SPUSŤ (scrDir + "/install_pkg.sh") s install_pkg.lst
                // install_pkg.sh nainstaluje balíčky uvedené v seznamu
         END IF
    END IF

    // -----------------------------
    // OBNOVENÍ KONFIGURAČNÍCH SOUBORŮ
    // -----------------------------
    IF (flg_Restore == 1) THEN
         VYPIŠ "restore banner"
         IF (není dry run a HYPRLAND_INSTANCE_SIGNATURE je nastaveno) THEN
              VYPOŘÍDAJ se s automatickým reloadem (hyprctl keyword misc:disable_autoreload)
         END IF

         // Spuštění obnovovacích skriptů:
         SPUSŤ (scrDir + "/restore_fnt.sh")
         SPUSŤ (scrDir + "/restore_cfg.sh")
         SPUSŤ (scrDir + "/restore_thm.sh")
         VYPIŠ, že se generují cache pro tapety
         IF (není dry run) THEN
              SPUSŤ "$HOME/.local/lib/hyde/swwwallcache.sh" -t ""
              SPUSŤ "$HOME/.local/lib/hyde/themeswitch.sh" -q  // Ignoruj chyby
              VYPIŠ informaci o reloadu Hyprland
         END IF
    END IF

    // -----------------------------
    // POST-INSTALAČNÍ KROK
    // -----------------------------
    IF (flg_Install == 1 AND flg_Restore == 1) THEN
         VYPIŠ "post-install banner"
         SPUSŤ (scrDir + "/install_pst.sh")
           // install_pst.sh provádí post-instalaci (např. čištění, další konfiguraci)
    END IF

    // -----------------------------
    // SPRÁVA SYSTEMD SLUŽEB
    // -----------------------------
    IF (flg_Service == 1) THEN
         VYPIŠ "service banner"
         Otevři soubor (scrDir + "/system_ctl.lst") a pro každý řádek (název služby):
              CHECK jestli služba již běží přes systemctl
              IF služba již existuje THEN
                   VYPIŠ, že služba je aktivní (skip)
              ELSE
                   VYPIŠ, že služba bude spuštěna
                   IF (není dry run) THEN
                        ENABLE a START službu přes systemctl
                   END IF
              END IF
         KONEC smyčky
    END IF

    // -----------------------------
    // Závěrečné hlášení a případný reboot
    // -----------------------------
    IF (flg_Install OR flg_Restore OR flg_Service jsou aktivní) THEN
         VYPIŠ dokončení instalace a umístění logů
         POŽÁDÁNÍ o reboot: "Je doporučeno restartovat systém. Chcete rebootovat? (y/N)"
         ČTENÍ vstupu
         IF (uživatel potvrdí reboot) THEN
              REBOOT systému pomocí systemctl
         ELSE
              VYPIŠ, že systém nebude restartován
         END IF
    END IF

END

