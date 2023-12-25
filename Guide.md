# Simple SetUp Guide

1. Get the latest release for your Operating System from [releases](https://github.com/yodaluca23/Warp-VPN-AdGuard-DNS/releases/latest).
2. Copy the path and remove any quotations from the copied path.
3. Open up CMD (Windows) or Terminal (MacOS & Linux)
4. Run ```[Path] register```.
5. (Optional) Run ```[Path] update -n [Whatever you want the device to be named, in your account.]```.
6. (Optional) Navigate to your user folder (Or system32, if you ran as administrator), and open a file called ```wgcf-account.toml```.
7. (Optional) Change whatever is in the quotations after ```license_key =``` to your current Warp License Key, and save the file.
8. Run ```[Path] generate```.
9. Your WireGuard ```wgcf-profile.conf``` file should now be in your user folder (Or system32, if you ran as administrator).
10. Transfer the WireGuard file to your device, and load it into any WireGuard Client app (I recommend [PassePartOut](https://passepartoutvpn.app)).
11. Your done! Now you can enable the VPN from your WireGuard Client, and be connected to CloudFlare's Warp VPN, while having no ads, using AdGuards DNS!
