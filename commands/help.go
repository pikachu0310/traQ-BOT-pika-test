package commands

import "example-bot/api"

func Help(ChannelID string, slice []string) {
	//message := "@BOT_pika_test /game"
	//OxGameHelp := ":blob_speedy_roll_inverse::blob_speedy_roll_inverse::blob_speedy_roll_inverse:早押しスタンプ:o::x:ゲーム Ver" + OxGameVersion + ":blob_speedy_roll::blob_speedy_roll::blob_speedy_roll:\n``@BOT_pika_test /game`` と入力することで遊べるよ！\n↓ ↓ 遊び方 ↓ ↓\n```\nBOTが3x3のマス上全てにランダムなスタンプを配置するので、\nマスと同じスタンプを押してマスを獲得し、一列揃えたら勝ち！(誰も揃わなかったら最も多かった人からランダム)\n```\n\nこのメッセージに:type_normal:を押すとノーマルモード\nこのメッセージに:crying-hard:を押すとハードモードで始まるよ！\n全9マスを埋めるTA(TimeAttack)モードもあるぞ！(↓のコマンドで出来る)(通常時でも全マスが埋まってたらTAモード扱いになる)\ntips:``/game start``,``/game start hard``,``/game ta``,``/game ta hard``でクイックスタート(この文章をスキップ)ができるよ！\nタイムが出るのでタイムアタックとしても楽しんで！ Enjoy! :party_blob:"
	message := "# :@BOT_pika_test:Ver" + OxGameVersion + " コマンド一覧\n## `@BOT_pika_test /game`\n-   早押しスタンプ:o::x:ゲーム で遊ぼう！\n#### `@BOT_pika_test /oisu`\n-   :oisu-1::oisu-2::oisu-3::oisu-4yoko:の順番を混ぜて:oisu-1::oisu-4yoko::oisu-3::oisu-2:するぜ\n#### `@BOT_pika_test /ping`\n-   pong!\n#### `@BOT_pika_test /help`\n-   これを表示する\n\nここから実験用(使用を想定していない)\n`@BOT_pika_test /slice <slice>`\n`@BOT_pika_test /stamp add <stampID>`\n`@BOT_pika_test /stamp remove <stampID>`\n`@BOT_pika_test /allstamps <num>`\n`@BOT_pika_test /edit <messageID> <content>`"
	api.PostMessage(ChannelID, message)
}
