# subcal (subnet calc)

サブネットの計算(10進数←→2進数)とかCIDRの計算をする

## できること

- 10進数、2進数相互変換
- IPリストの一覧出力
- ネットワーク部とホスト部の色分け

## 使い方

```bash
$ subcal 126.227.243.233/24
IPv4	CIDR	IPv4(Bin)	Mask
126.227.243.233	24	01111110111000111111001111101001	11111111111111111111111100000000
```

色付き出力

```bash
$ subcal -c 126.227.243.233/24
IPv4	CIDR	IPv4(Bin)	Mask
126.227.243.233	24	01111110111000111111001111101001	11111111111111111111111100000000
```
