# 📌 HyprZen – A Zen-like Installer & Configurator for Hyprland on Arch Linux

🚧 **HyprZen is in active development!** 🚧  
The goal is to create a **minimalist, Go-powered installer** for **Hyprland on Arch Linux**, fully replacing complex shell scripts with a **clean, interactive TUI**.


---

## 🔧 System Requirements

- Must be running **Arch Linux** (a minimal installation, not the archinstall variant)  
- **GRUB** bootloader is required  
- **git**
- **go**

---

## ✨ Planned Features

✅ **One-command installation** of Hyprland & essential packages  
✅ **Enable system services** automatically (`systemctl enable`)  
✅ **AUR support** (optional) for extended customization  
✅ **Dotfiles & themes restoration** for a seamless environment  
✅ **Nvidia support toggle** (optional)  
✅ **Zen-like simplicity** – no clutter, just what you need  

---

## 🔧 Project Status

HyprZen is currently a **work in progress**. Contributions, feedback, and ideas are welcome! 🚀

---

## 💡 Why HyprZen?

❌ **No shell scripts** – only Go  
❌ **No unnecessary features** – just what’s needed  
❌ **No forced configurations** – full user control  

---

## 🚀 Coming Soon

- **TUI-based selection & logging**
- **Better package management & customization**
- **Pre-built binaries for easy installation**  

Stay tuned!  

---

## 🛠 How to Run (Development Mode)

```sh
git clone https://github.com/tisonpatrik/HyprZen.git
cd HyprZen
make build
./build/main

