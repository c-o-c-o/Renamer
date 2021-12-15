# Renamer

音声合成ファイルの出力ファイルを元にファイルをリネームするアプリケーションです。  
ついでにテキストファイルの中も整形します。

# 使い方

```
renamer.exe [options] [files]
options
  -t テキストファイルのパス

files
  リネームするファイルのパス
```

```
example
  renamer.exe -t セリフ.txt セリフ.wav
```

# 設定

VOICEROID2 と VOICEBOX 用に設定しています。  
一部のキャラクターは名前を置換するように設定しています。  
詳細は `setting.yml` を参照してください。

#

This software is released under the MIT License, see LICENSE.
