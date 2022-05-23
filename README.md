# sudo-implementation
sudoの簡易的な実装

## 使用手順

```
go build get_id.go
go build camp_sudo.go
sudo chown root:root camp_sudo # 所有者をrootにする
sudo chmod u+s camp_sudo # suidビットをセットする
sudo touch /etc/camp_sudoers #空のsudoersファイルを作る
./camp_sudo -cmd /usr/sbin/ifconfig #sudoersにユーザー名が無いのでエラー
sudo vim /etc/camp_sudoers #sudoersにユーザー名を書き込む
./camp_sudo /usr/sbin/ifconfig
./camp_sudo -cmd ./get_id
./camp_sudo -u www-data -cmd ./get_id
```