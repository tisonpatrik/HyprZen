BEGIN
    // Získání adresáře, kde se skript nachází
    SET scrDir = adresář, kde se nachází tento skript

    // Načtení globálních funkcí a proměnných z global_fn.sh
    SOURCE (scrDir + "/global_fn.sh")
      // global_fn.sh poskytuje funkce jako print_log, prompt_timer, nvidia_detect, atd.

    // Načtení režimu suchého běhu (dry run), pokud není nastavena, nastav na 0
    SET flg_DryRun = (hodnota z prostředí, nebo výchozí 0)

    // -----------------------------------------------------------
    // Konfigurace GRUB bootloaderu
    // -----------------------------------------------------------
    IF (balíček "grub" je nainstalován AND soubor /boot/grub/grub.cfg existuje) THEN
        LOG "GRUB bootloader detekován"

        // Kontrola, zda již nebyly vytvořeny zálohy konfigurací GRUB
        IF (záloha /etc/default/grub.hyde.bkp a /boot/grub/grub.hyde.bkp neexistují) THEN
            LOG "Zálohování a konfigurace GRUB"
            // Vytvoření záloh konfiguračních souborů
            COPY /etc/default/grub -> /etc/default/grub.hyde.bkp
            COPY /boot/grub/grub.cfg -> /boot/grub/grub.hyde.bkp

            // Pokud je detekována Nvidia GPU, přidat volbu nvidia_drm.modeset=1
            IF (nvidia_detect vrací true) THEN
                LOG "Nvidia detekována: přidání nvidia_drm.modeset=1"
                // Získat aktuální GRUB_CMDLINE_LINUX_DEFAULT a odstranit případné duplicitní volby
                UPDATE hodnotu GRUB_CMDLINE_LINUX_DEFAULT v /etc/default/grub, přidat 'nvidia_drm.modeset=1'
            END IF

            // Výzva uživateli k výběru GRUB tématu
            LOG "Výběr tématu GRUB"
            DISPLAY možnosti: [1] Retroboot (tmavé), [2] Pochita (světlé), nebo vynechat výběr
            READ uživatelský vstup (grubopt)
            SWITCH grubopt:
                CASE 1: SET grubtheme = "Retroboot"
                CASE 2: SET grubtheme = "Pochita"
                DEFAULT: SET grubtheme = "None"
            END SWITCH

            // Nastavení tématu nebo jeho přeskočení
            IF (grubtheme == "None") THEN
                LOG "Téma GRUB se přeskočí"
                MODIFIKUJ /etc/default/grub tak, aby bylo vypnuto nastavení tématu
            ELSE
                LOG "Nastavení tématu GRUB: " + grubtheme
                // Rozbalit tar.gz archiv tématu do adresáře GRUB témat
                EXTRACT (cloneDir + "/Source/arcs/Grub_" + grubtheme + ".tar.gz") -> /usr/share/grub/themes/
                // Aktualizovat /etc/default/grub s novými nastaveními (GRUB_DEFAULT, GRUB_GFXMODE, GRUB_THEME atd.)
                UPDATE /etc/default/grub s odpovídajícími řádky
            END IF

            // Vygenerovat nový konfigurační soubor GRUB
            RUN sudo grub-mkconfig -o /boot/grub/grub.cfg
        ELSE
            LOG "GRUB již byl nakonfigurován, přeskočeno..."
        END IF
    END IF

    // -----------------------------------------------------------
    // Konfigurace systemd-boot
    // -----------------------------------------------------------
    IF (balíček "systemd" je nainstalován AND nvidia_detect je true AND bootctl status indikuje systemd-boot) THEN
        LOG "systemd-boot detekován"

        // Kontrola záloh konfigurací systemd-boot
        IF (počet záloh (*.conf.hyde.bkp) není stejný jako počet *.conf souborů) THEN
            LOG "Konfigurace systemd-boot s přidáním nvidia_drm.modeset=1"
            FOR EACH konfigurační soubor (*.conf) v /boot/loader/entries/ DO
                BACKUP soubor -> soubor.hyde.bkp
                // Odstranit nepotřebné volby a přidat volbu nvidia_drm.modeset=1 spolu s 'quiet splash'
                UPDATE řádek začínající "options" v souboru
            END FOR
        ELSE
            LOG "systemd-boot již byl nakonfigurován, přeskočeno..."
        END IF
    END IF

    // -----------------------------------------------------------
    // Úprava konfigurace pacman
    // -----------------------------------------------------------
    IF (/etc/pacman.conf existuje AND záloha /etc/pacman.conf.hyde.bkp neexistuje) THEN
        LOG "Úprava konfigurace pacman"
        IF (flg_DryRun není aktivní) THEN
            // Vytvoření zálohy pacman.conf
            COPY /etc/pacman.conf -> /etc/pacman.conf.hyde.bkp
            // Přidání voleb: Color, ILoveCandy, VerbosePkgLists, ParallelDownloads a odkomentování multilib
            MODIFIKUJ /etc/pacman.conf podle požadavků
            // Aktualizace systému
            RUN sudo pacman -Syyu
            RUN sudo pacman -Fy
        END IF
    ELSE
        LOG "Pacman již byl nakonfigurován, přeskočeno..."
    END IF

    // -----------------------------------------------------------
    // Přidání Chaotic AUR repozitáře
    // -----------------------------------------------------------
    IF (soubor /etc/pacman.conf obsahuje "[chaotic-aur]") THEN
        LOG "Chaotic AUR již existuje, přeskočeno..."
    ELSE
        // Výzva k instalaci Chaotic AUR s timeoutem
        PROMPT "Chcete nainstalovat Chaotic AUR? [y/n] | q k ukončení" s timeoutem 120 sekund
        READ uživatelský vstup (PROMPT_INPUT)
        SWITCH PROMPT_INPUT:
            CASE 'y' or 'Y': SET is_chaotic_aur = true
            CASE 'n' or 'N': SET is_chaotic_aur = false
            CASE 'q' or 'Q': LOG "Ukončuji" a EXIT
            DEFAULT: SET is_chaotic_aur = true
        END SWITCH
        IF (is_chaotic_aur == true) THEN
            // Inicializace pacman klíče a spuštění skriptu pro Chaotic AUR
            RUN sudo pacman-key --init
            RUN (scrDir + "/chaotic_aur.sh") s parametrem --install
        END IF
    END IF

END

