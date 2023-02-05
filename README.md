# traQ-BOT-pika-test
traQの @BOT_pika-test の全て
# pika_test

## 僕が公開シェルになった事件
##### [#gps/times/pikachu0310/Botのログを解説する会MD](https://md.trap.jp/qLi5fWPpTqixaBaODkKQfw)によくまとまってます！
![image.png](https://wiki.trap.jp/files/63dfe055c50373001473b24a)
![image.png](https://wiki.trap.jp/files/63dfe07ec50373001473b24c)
現場: https://q.trap.jp/messages/57e4a5c6-90d5-4ea2-ac81-51558648fe29

皆がどんなことをやろうとしたのかという[講習会が開かれました](https://q.trap.jp/channels/event/workshop/pika_test/jikkyo)(:@oribe:さんに感謝)

## その後の殺害事件
[悪魔による殺害現場](https://q.trap.jp/messages/6ba46260-f33c-4f40-b4a0-81e284e8e669)
![image.png](https://wiki.trap.jp/files/63dfdfddc50373001473b249)
[:@kegra:さんによるえげつある漫画](https://q.trap.jp/messages/0cd39714-5e9f-4bf3-a218-440fe168af3a)
![image.png](https://wiki.trap.jp/files/63dfde61c50373001473b244)
![image.png](https://wiki.trap.jp/files/63dfde67c50373001473b245)
![image.png](https://wiki.trap.jp/files/63dfde8ec50373001473b248)
![image.png](https://wiki.trap.jp/files/63dfde74c50373001473b247)

# 出来る事
### 早見表 (大体これがすべて)
| コマンド | 使い方                                      | 例                                      | 何をするか                               |
| -------- | ------------------------------------------- | --------------------------------------- | ---------------------------------------- |
| /game    | /game (start)                               | /game start                             | 早押しスタンプ:o::x:ゲームを開始するぞ！ |
| /tag     | /tag (UserID) (year)                        | /tag pikachu 2022                       | ある年につけられたタグ一覧を表示するぞ！ |
| /sql     | /sql [SQL文]                                | /sql select title, content from test3； | 任意のsql文を実行するぞ！                |
| :mag:    | :mag:or:Internet_Explorer: [検索キーワード] | :mag:VRChat お砂糖 意味                 | I'm feeling luckyするぞ！                |
| /info    | /info [人物名]                              | /info :@pikachu:                        | ユーザーかスタンプの詳細を表示するぞ！   |
| join     |                             メンション必須                |  @BOT_pika_test join                                       |                                    BOTがチャンネルに参加する      |
| leave    |           メンション必須                                  |                            @BOT_pika_test leave            |               BOTがチャンネルから抜ける                           |
| help | /help | /help | これを表示するぞ！ |


## 詳細
### /game (早押しスタンプゲーム)
![image.png](https://wiki.trap.jp/files/63dff8aac50373001473b251)
遊び方 : BOTが3x3のマス上全てにランダムなスタンプを配置するので、
マスと同じスタンプを押してマスを獲得し、一列揃えたら勝ち！(誰も揃わなかったら最も多かった人からランダム)
スタンプの上にカーソルをホバーさせてスタンプ名を見るのはOKだぞ！

現状normalは:@shobon:君の[5.152秒が最速記録](https://q.trap.jp/messages/87045a1c-9cf8-48ac-b18c-e982cb680fab)
1列揃えば終わるnormalモードと、全9マスが埋まり終わるまで終わらないtimeAttackモードがある。
また、通常のスタンプのnormalモードとスタンプにエフェクトが5個付くhardモードがある。
なので、タイムアタックとしては、normal、timeAttack、hard、timeAttackHardの4種類が存在する。
timeAttackに関しては参加した人数も記録されるので、平等な条件で比較することもできる。
是非タイムアタックとしても、バトルとしても楽しんで貰いたい！(深夜にやるとめっちゃ面白い系の奴なので)

- /game
- /game start
- /game start hard
- /game timeattack
- /game ta
- /game timeattack hard
- /game ta hard
- /game reset
- /gane debug

### /tag (タグ一覧を表示)
![image.png](https://wiki.trap.jp/files/63dff8fac50373001473b252)

今年につけられたタグを振り返ろう！ということで作った機能。タグが付けられた日時も分かるので、タグで1年を振り返ってみよう！
- /tag
- /tag pikachu
- /tag pikachu 2022

### /sql (公開bash)
![image.png](https://wiki.trap.jp/files/63dff92ec50373001473b253)
データベースの操作で使うsqlを任意実行できる機能。何やら悪さができるらしいぞ...？
- /sql [sql文]
- /sql system [bash文]
- /docker up
- /docker down
- /docker restart

### mag (I'm feeling luckyする)
![image.png](https://wiki.trap.jp/files/63dff99bc50373001473b254)
googleの検索結果の一番上のURLを出力します。vsきなの
- :mag:あいうえお かきく
- あい:mag:うえお かきく
- あいうえお:mag: かきく
- あい:mag_right:うえお かきく
- あい:Internet_Explorer:うえお かきく

### info (詳細を表示)
![image.png](https://wiki.trap.jp/files/63dffa2fc50373001473b255)
ユーザーまたはスタンプの詳細を出力する。UUIDをサッと知りたいときに便利かも
- /info pikachu
- /info :pikachu:
- /info :@pikachu:

### join (チャンネルに参加)
- @BOT_pika_test join
- (join|いらっしゃい|oisu-|:oisu-1::oisu-2::oisu-3::oisu-4yoko:|おいすー) のいずれかなら反応する

### leave (チャンネルから抜ける)
- @BOT_pika_test leave
- (leave|さようなら|:wave:|ばいばい) のいずれかなら反応する

#### ここから古いhelpからの引用(あまり使わないコマンド)
#### `@BOT_pika_test /oisu`
-   :oisu-1::oisu-2::oisu-3::oisu-4yoko:の順番を混ぜて:oisu-1::oisu-4yoko::oisu-3::oisu-2:するぞ！
#### `@BOT_pika_test /ping`
-   pong! するぞ！
#### `@BOT_pika_test /help`
-   これを表示するぞ！

ここから実験用(使用を想定していない)
`@BOT_pika_test /slice <slice>`
`@BOT_pika_test /stamp add <stampID>`
`@BOT_pika_test /stamp remove <stampID>`
`@BOT_pika_test /allstamps <num>`
`@BOT_pika_test /edit <messageID> <content>`

### 早見表 (再掲 大体これがすべて)
| コマンド | 使い方                                      | 例                                      | 何をするか                               |
| -------- | ------------------------------------------- | --------------------------------------- | ---------------------------------------- |
| /game    | /game (start)                               | /game start                             | 早押しスタンプ:o::x:ゲームを開始するぞ！ |
| /tag     | /tag (UserID) (year)                        | /tag pikachu 2022                       | ある年につけられたタグ一覧を表示するぞ！ |
| /sql     | /sql [SQL文]                                | /sql select title, content from test3； | 任意のsql文を実行するぞ！                |
| :mag:    | :mag:or:Internet_Explorer: [検索キーワード] | :mag:VRChat お砂糖 意味                 | I'm feeling luckyするぞ！                |
| /info    | /info [人物名]                              | /info :@pikachu:                        | ユーザーかスタンプの詳細を表示するぞ！   |
| join     |                             メンション必須                |  @BOT_pika_test join                                       |                                    BOTがチャンネルに参加する      |
| leave    |           メンション必須                                  |                            @BOT_pika_test leave            |               BOTがチャンネルから抜ける                           |
| help | /help | /help | これを表示するぞ！ |
