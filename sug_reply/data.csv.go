package sug_reply

import (
	"code.byted.org/gopkg/jsonx"
	"fmt"
	"strings"
)

type Reply struct {
	ReplyText   string `json:"reply_text"`
	ReplyEmoji  string `json:"reply_emoji"`
	StarlingKey string `json:"starling_key"`
}

var rawData = `
{
  "Scripted Comedy": [
    {"reply_text": "HAHA", "reply_emoji": "😂", "starling_key": "dm_suggestreply_haha"},
    {"reply_text": "I'm crying", "reply_emoji": "🤣", "starling_key": "dm_suggestreply_im_crying"},
    {"reply_text": "LMAO", "reply_emoji": "😹", "starling_key": "dm_suggestreply_lmao"},
    {"reply_text": "Comedy gold!", "reply_emoji": "😆", "starling_key": "dm_suggestreply_comedy_gold"},
    {"reply_text": "People are hilarious", "reply_emoji": "😅", "starling_key": "dm_suggestreply_people_hilar"}
  ],
  "Pranks": [
    {"reply_text": "LOL", "reply_emoji": "😂", "starling_key": "dm_suggestreply_lol"},
    {"reply_text": "Made my day", "reply_emoji": "😆", "starling_key": "dm_suggestreply_made_my_day"},
    {"reply_text": "So funny!", "reply_emoji": "🤣", "starling_key": "dm_suggestreply_so_funny"},
    {"reply_text": "Epic prank!", "reply_emoji": "😅", "starling_key": "dm_suggestreply_epic_prank"},
    {"reply_text": "OMG", "reply_emoji": "😹", "starling_key": "dm_suggestreply_omg"}
  ],
  "Hilarious Fails": [
    {"reply_text": "I can't", "reply_emoji": "😅", "starling_key": "dm_suggestreply_cant"},
    {"reply_text": "Made my day", "reply_emoji": "😂", "starling_key": "dm_suggestreply_made_my_day2"},
    {"reply_text": "I'm screaming", "reply_emoji": "🤣", "starling_key": "dm_suggestreply_im_screaming"},
    {"reply_text": "Epic fail", "reply_emoji": "😹", "starling_key": "dm_suggestreply_epic_fail"},
    {"reply_text": "Wheezing!", "reply_emoji": "😆", "starling_key": "dm_suggestreply_wheezing"}
  ],
  "Other Comedy": [
    {"reply_text": "Too funny!", "reply_emoji": "😹", "starling_key": "dm_suggestreply_too_funny"},
    {"reply_text": "HAHA", "reply_emoji": "😂", "starling_key": "dm_suggestreply_haha"},
    {"reply_text": "Comedy gold!", "reply_emoji": "🤣", "starling_key": "dm_suggestreply_comedy_gold2"},
    {"reply_text": "I love TikTok", "reply_emoji": "😂", "starling_key": "dm_suggestreply_love_tiktok"},
    {"reply_text": "Too good!", "reply_emoji": "😅", "starling_key": "dm_suggestreply_too_good"}
  ],
  "Lip-sync": [
    {"reply_text": "Too good!", "reply_emoji": "💯", "starling_key": "dm_suggestreply_too_good1"},
    {"reply_text": "Huge fan!", "reply_emoji": "😲", "starling_key": "dm_suggestreply_huge_fan"},
    {"reply_text": "Wow", "reply_emoji": "👏", "starling_key": "dm_suggestreply_wow"},
    {"reply_text": "Iconic", "reply_emoji": "👑", "starling_key": "dm_suggestreply_iconic"},
    {"reply_text": "So talented!", "reply_emoji": "👏", "starling_key": "dm_suggestreply_so_talented"}
  ],
  "Scripted Drama": [
    {"reply_text": "Viral material!", "reply_emoji": "😲", "starling_key": "dm_suggestreply_viral_material"},
    {"reply_text": "That's so good!", "reply_emoji": "😆", "starling_key": "dm_suggestreply_thats_so_good"},
    {"reply_text": "I'm speechless", "reply_emoji": "😶", "starling_key": "dm_suggestreply_im_speechless"},
    {"reply_text": "Cringe", "reply_emoji": "😬", "starling_key": "dm_suggestreply_cringe"},
    {"reply_text": "Wait a sec", "reply_emoji": "👀", "starling_key": "dm_suggestreply_wait_a_sec"}
  ],
  "Dance": [
    {"reply_text": "Too good!", "reply_emoji": "👏", "starling_key": "dm_suggestreply_too_good2"},
    {"reply_text": "Perfection", "reply_emoji": "💯", "starling_key": "dm_suggestreply_perfection"},
    {"reply_text": "Cringe", "reply_emoji": "😬", "starling_key": "dm_suggestreply_cringe"},
    {"reply_text": "Amazing", "reply_emoji": "😍", "starling_key": "dm_suggestreply_amazing"},
    {"reply_text": "Love it!", "reply_emoji": "💃", "starling_key": "dm_suggestreply_love_it"}
  ],
  "Finger Dance & Basic Dance": [
    {"reply_text": "Vibe check", "reply_emoji": "✅", "starling_key": "dm_suggestreply_vibe_check"},
    {"reply_text": "Wow!", "reply_emoji": "👏", "starling_key": "dm_suggestreply_wow2"},
    {"reply_text": "Mind-blowing!", "reply_emoji": "🤯", "starling_key": "dm_suggestreply_mind_blowing"},
    {"reply_text": "So detailed", "reply_emoji": "🔍", "starling_key": "dm_suggestreply_so_detailed"},
    {"reply_text": "No way!", "reply_emoji": "😳", "starling_key": "dm_suggestreply_no_way"}
  ],
  "Singing & Instruments": [
    {"reply_text": "Huge fan", "reply_emoji": "😲", "starling_key": "dm_suggestreply_huge_fan1"},
    {"reply_text": "Too good!", "reply_emoji": "💯", "starling_key": "dm_suggestreply_too_good1"},
    {"reply_text": "So talented!", "reply_emoji": "👏", "starling_key": "dm_suggestreply_so_talented"},
    {"reply_text": "That's incredible", "reply_emoji": "🤩", "starling_key": "dm_suggestreply_thats_incredib"},
    {"reply_text": "Perfection", "reply_emoji": "👌", "starling_key": "dm_suggestreply_perfection2"}
  ],
  "Graphic Art": [
    {"reply_text": "Amazing", "reply_emoji": "😍", "starling_key": "dm_suggestreply_amazing"},
    {"reply_text": "Unreal!", "reply_emoji": "😲", "starling_key": "dm_suggestreply_unreal"},
    {"reply_text": "Perfection", "reply_emoji": "💯", "starling_key": "dm_suggestreply_perfection"},
    {"reply_text": "So creative!", "reply_emoji": "🎨", "starling_key": "dm_suggestreply_so_creative"},
    {"reply_text": "That's incredible", "reply_emoji": "🤩", "starling_key": "dm_suggestreply_thats_incredib"}
  ],
  "Other Art": [
    {"reply_text": "So creative!", "reply_emoji": "🎨", "starling_key": "dm_suggestreply_so_creative"},
    {"reply_text": "Beautiful", "reply_emoji": "🌟", "starling_key": "dm_suggestreply_beautiful"},
    {"reply_text": "So detailed", "reply_emoji": "🔍", "starling_key": "dm_suggestreply_so_detailed"},
    {"reply_text": "Stunning!", "reply_emoji": "😲", "starling_key": "dm_suggestreply_stunning"},
    {"reply_text": "That's unbelievable", "reply_emoji": "😱", "starling_key": "dm_suggestreply_thats_unbeliev"}
  ],
  "Magic": [
    {"reply_text": "I can't believe it", "reply_emoji": "😲", "starling_key": "dm_suggestreply_cant_believe_it"},
    {"reply_text": "Mind-boggling...", "reply_emoji": "🤔", "starling_key": "dm_suggestreply_mind_boggling"},
    {"reply_text": "No way!", "reply_emoji": "😮", "starling_key": "dm_suggestreply_no_way2"},
    {"reply_text": "Unreal!", "reply_emoji": "🤯", "starling_key": "dm_suggestreply_unreal2"},
    {"reply_text": "I'm shook", "reply_emoji": "😱", "starling_key": "dm_suggestreply_im_shook"}
  ],
  "Professional Special Effects": [
    {"reply_text": "Whoa!", "reply_emoji": "😮", "starling_key": "dm_suggestreply_whoa"},
    {"reply_text": "Mind-boggling...", "reply_emoji": "🤔", "starling_key": "dm_suggestreply_mind_boggling"},
    {"reply_text": "I can't believe it", "reply_emoji": "😳", "starling_key": "dm_suggestreply_cant_believe_i"},
    {"reply_text": "Mindblown!", "reply_emoji": "🤯", "starling_key": "dm_suggestreply_mindblown"},
    {"reply_text": "I'm shook", "reply_emoji": "😱", "starling_key": "dm_suggestreply_im_shook"}
  ],
  "DIY & Handicrafts": [
    {"reply_text": "So creative!", "reply_emoji": "🎨", "starling_key": "dm_suggestreply_so_creative"},
    {"reply_text": "So cool!", "reply_emoji": "🤩", "starling_key": "dm_suggestreply_so_cool"},
    {"reply_text": "Incredible", "reply_emoji": "😯", "starling_key": "dm_suggestreply_incredible"},
    {"reply_text": "Perfection", "reply_emoji": "💯", "starling_key": "dm_suggestreply_perfection"},
    {"reply_text": "I want one", "reply_emoji": "😍", "starling_key": "dm_suggestreply_i_want_one"}
  ],
  "Other Talent": [
    {"reply_text": "Too good!", "reply_emoji": "💯", "starling_key": "dm_suggestreply_too_good1"},
    {"reply_text": "Impressive", "reply_emoji": "🤩", "starling_key": "dm_suggestreply_impressive"},
    {"reply_text": "So talented!", "reply_emoji": "🔥", "starling_key": "dm_suggestreply_so_talented2"},
    {"reply_text": "Whoa!", "reply_emoji": "😮", "starling_key": "dm_suggestreply_whoa"},
    {"reply_text": "Mind-boggling...", "reply_emoji": "🤔", "starling_key": "dm_suggestreply_mind_boggling"}
  ],
  "Romance": [
    {"reply_text": "Congrats!", "reply_emoji": "😭", "starling_key": "dm_suggestreply_congrats"},
    {"reply_text": "Goals", "reply_emoji": "😍", "starling_key": "dm_suggestreply_goals"},
    {"reply_text": "So cute!", "reply_emoji": "🥰", "starling_key": "dm_suggestreply_so_cute"},
    {"reply_text": "Aww", "reply_emoji": "😍", "starling_key": "dm_suggestreply_aww"},
    {"reply_text": "Cringe", "reply_emoji": "😬", "starling_key": "dm_suggestreply_cringe"}
  ],
  "Family": [
    {"reply_text": "Not me crying", "reply_emoji": "😓", "starling_key": "dm_suggestreply_not_me_crying"},
    {"reply_text": "Happy tears", "reply_emoji": "🥹", "starling_key": "dm_suggestreply_happy_tears"},
    {"reply_text": "Are you serious", "reply_emoji": "😢", "starling_key": "dm_suggestreply_are_you_serious"},
    {"reply_text": "Tears are falling", "reply_emoji": "😭", "starling_key": "dm_suggestreply_tears_are_falling"},
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"}
  ],
  "Babies": [
    {"reply_text": "Aww", "reply_emoji": "😍", "starling_key": "dm_suggestreply_aww"},
    {"reply_text": "So cute!", "reply_emoji": "🥰", "starling_key": "dm_suggestreply_so_cute"},
    {"reply_text": "LOL", "reply_emoji": "😂", "starling_key": "dm_suggestreply_lol"},
    {"reply_text": "That's so silly", "reply_emoji": "😆", "starling_key": "dm_suggestreply_thats_so_silly"},
    {"reply_text": "So adorable!", "reply_emoji": "😘", "starling_key": "dm_suggestreply_so_adorable"}
  ],
  "Pets": [
    {"reply_text": "OMG!", "reply_emoji": "😲", "starling_key": "dm_suggestreply_omg2"},
    {"reply_text": "So cute!", "reply_emoji": "🥰", "starling_key": "dm_suggestreply_so_cute"},
    {"reply_text": "I want one", "reply_emoji": "😍", "starling_key": "dm_suggestreply_i_want_one"},
    {"reply_text": "I can't believe it!", "reply_emoji": "🙀", "starling_key": "dm_suggestreply_i_cant_believe_it"},
    {"reply_text": "Sadness", "reply_emoji": "😢", "starling_key": "dm_suggestreply_sadness"}
  ],
  "Farm Animals": [
    {"reply_text": "That's wild", "reply_emoji": "🤠", "starling_key": "dm_suggestreply_thats_wild"},
    {"reply_text": "OMG!", "reply_emoji": "😲", "starling_key": "dm_suggestreply_omg2"},
    {"reply_text": "I can't", "reply_emoji": "🤣", "starling_key": "dm_suggestreply_i_cant"},
    {"reply_text": "Too funny!", "reply_emoji": "😹", "starling_key": "dm_suggestreply_too_funny"},
    {"reply_text": "Love animals", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_animals"}
  ],
  "Other Animals": [
    {"reply_text": "Whoa!", "reply_emoji": "😮", "starling_key": "dm_suggestreply_whoa"},
    {"reply_text": "Too funny!", "reply_emoji": "😹", "starling_key": "dm_suggestreply_too_funny"},
    {"reply_text": "Are you serious", "reply_emoji": "😢", "starling_key": "dm_suggestreply_are_you_serious"},
    {"reply_text": "That's wild", "reply_emoji": "🤠", "starling_key": "dm_suggestreply_thats_wild"},
    {"reply_text": "Love animals", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_animals"}
  ],
  "Scenery & Plants": [
    {"reply_text": "I'm speechless", "reply_emoji": "😶", "starling_key": "dm_suggestreply_im_speechless"},
    {"reply_text": "Mind-boggling...", "reply_emoji": "🤔", "starling_key": "dm_suggestreply_mind_boggling"},
    {"reply_text": "Incredible!", "reply_emoji": "😯", "starling_key": "dm_suggestreply_incredible2"},
    {"reply_text": "Out of this world", "reply_emoji": "👽", "starling_key": "dm_suggestreply_out_of_this_world"},
    {"reply_text": "Perfection", "reply_emoji": "💯", "starling_key": "dm_suggestreply_perfection"}
  ],
  "Food Display": [
    {"reply_text": "Yums", "reply_emoji": "😋", "starling_key": "dm_suggestreply_yums"},
    {"reply_text": "Looks tasty!", "reply_emoji": "😍", "starling_key": "dm_suggestreply_looks_tasty"},
    {"reply_text": "I need!", "reply_emoji": "😋", "starling_key": "dm_suggestreply_i_need"},
    {"reply_text": "Drooling", "reply_emoji": "🤤", "starling_key": "dm_suggestreply_drooling"},
    {"reply_text": "Unreal!", "reply_emoji": "🤯", "starling_key": "dm_suggestreply_unreal2"}
  ],
  "Cooking": [
    {"reply_text": "Yums", "reply_emoji": "😋", "starling_key": "dm_suggestreply_yums"},
    {"reply_text": "Looks tasty!", "reply_emoji": "😍", "starling_key": "dm_suggestreply_looks_tasty"},
    {"reply_text": "I need", "reply_emoji": "😋", "starling_key": "dm_suggestreply_i_need1"},
    {"reply_text": "Drooling", "reply_emoji": "🤤", "starling_key": "dm_suggestreply_drooling"},
    {"reply_text": "Unreal!", "reply_emoji": "🤯", "starling_key": "dm_suggestreply_unreal2"}
  ],
  "Drinks": [
    {"reply_text": "I want", "reply_emoji": "😍", "starling_key": "dm_suggestreply_i_want"},
    {"reply_text": "Cheers!", "reply_emoji": "🥂", "starling_key": "dm_suggestreply_cheers"},
    {"reply_text": "Looks bussin", "reply_emoji": "😲", "starling_key": "dm_suggestreply_looks_bussin"},
    {"reply_text": "Vibe check", "reply_emoji": "✅", "starling_key": "dm_suggestreply_vibe_check"},
    {"reply_text": "Too good!", "reply_emoji": "💯", "starling_key": "dm_suggestreply_too_good1"}
  ],
  "Food Tour & Recommendations": [
    {"reply_text": "Looks great", "reply_emoji": "😲", "starling_key": "dm_suggestreply_looks_great"},
    {"reply_text": "I want!", "reply_emoji": "😍", "starling_key": "dm_suggestreply_i_want2"},
    {"reply_text": "Vibe check", "reply_emoji": "✅", "starling_key": "dm_suggestreply_vibe_check"},
    {"reply_text": "Looks tasty!", "reply_emoji": "😍", "starling_key": "dm_suggestreply_looks_tasty"},
    {"reply_text": "Yums", "reply_emoji": "😋", "starling_key": "dm_suggestreply_yums"}
  ],
  "Mukbangs & Tasting": [
    {"reply_text": "I want", "reply_emoji": "😍", "starling_key": "dm_suggestreply_i_want"},
    {"reply_text": "OMG!", "reply_emoji": "😲", "starling_key": "dm_suggestreply_omg2"},
    {"reply_text": "Yums!", "reply_emoji": "😋", "starling_key": "dm_suggestreply_yums2"},
    {"reply_text": "This was healing", "reply_emoji": "💆‍♀️", "starling_key": "dm_suggestreply_this_was_healing"},
    {"reply_text": "I want", "reply_emoji": "😍", "starling_key": "dm_suggestreply_i_want"}
  ],
  "Travel": [
    {"reply_text": "Manifesting this", "reply_emoji": "✨", "starling_key": "dm_suggestreply_manifesting_this"},
    {"reply_text": "I wanna go!", "reply_emoji": "🌎", "starling_key": "dm_suggestreply_i_wanna_go"},
    {"reply_text": "Mind-blowing!", "reply_emoji": "🤯", "starling_key": "dm_suggestreply_mind_blowing"},
    {"reply_text": "Travel goals", "reply_emoji": "💯", "starling_key": "dm_suggestreply_travel_goals"},
    {"reply_text": "When's it my turn", "reply_emoji": "😭", "starling_key": "dm_suggestreply_whens_it_my_turn"}
  ],
  "Diary & VLOG": [
    {"reply_text": "Quality content", "reply_emoji": "💯", "starling_key": "dm_suggestreply_quality_content"},
    {"reply_text": "Love this!", "reply_emoji": "🎥", "starling_key": "dm_suggestreply_love_this1"},
    {"reply_text": "Favs", "reply_emoji": "😍", "starling_key": "dm_suggestreply_favs"},
    {"reply_text": "So cool!", "reply_emoji": "😎", "starling_key": "dm_suggestreply_so_cool1"},
    {"reply_text": "I'm shook", "reply_emoji": "😱", "starling_key": "dm_suggestreply_im_shook"}
  ],
  "Recreation Facilities": [
    {"reply_text": "Beautiful", "reply_emoji": "🌟", "starling_key": "dm_suggestreply_beautiful"},
    {"reply_text": "Stay fit", "reply_emoji": "💪", "starling_key": "dm_suggestreply_stay_fit"},
    {"reply_text": "Whoa!", "reply_emoji": "😮", "starling_key": "dm_suggestreply_whoa"},
    {"reply_text": "I wanna go!", "reply_emoji": "🤩", "starling_key": "dm_suggestreply_i_wanna_go2"},
    {"reply_text": "That's awesome", "reply_emoji": "😁", "starling_key": "dm_suggestreply_thats_awesome"}
  ],
  "Campus Life": [
    {"reply_text": "Brings back memories", "reply_emoji": "👨‍🎓", "starling_key": "dm_suggestreply_brings_back_memories"},
    {"reply_text": "Wait a sec", "reply_emoji": "👀", "starling_key": "dm_suggestreply_wait_a_sec"},
    {"reply_text": "LOL", "reply_emoji": "😂", "starling_key": "dm_suggestreply_lol"},
    {"reply_text": "Vibe check", "reply_emoji": "✅", "starling_key": "dm_suggestreply_vibe_check"},
    {"reply_text": "Whoa!", "reply_emoji": "😮", "starling_key": "dm_suggestreply_whoa"}
  ],
  "Work & Jobs": [
    {"reply_text": "HAHA", "reply_emoji": "😂", "starling_key": "dm_suggestreply_haha"},
    {"reply_text": "Interesting", "reply_emoji": "🤔", "starling_key": "dm_suggestreply_interesting"},
    {"reply_text": "Mindblown!", "reply_emoji": "🤯", "starling_key": "dm_suggestreply_mindblown"},
    {"reply_text": "That's impressive", "reply_emoji": "😱", "starling_key": "dm_suggestreply_thats_impressive"},
    {"reply_text": "Same", "reply_emoji": "🙃", "starling_key": "dm_suggestreply_same"}
  ],
  "Life Hacks": [
    {"reply_text": "Saving this", "reply_emoji": "🤝", "starling_key": "dm_suggestreply_saving_this"},
    {"reply_text": "So cool", "reply_emoji": "👍", "starling_key": "dm_suggestreply_so_cool5"},
    {"reply_text": "I need", "reply_emoji": "😮", "starling_key": "dm_suggestreply_I_need3"},
    {"reply_text": "Wait a sec", "reply_emoji": "👀", "starling_key": "dm_suggestreply_wait_a_sec"},
    {"reply_text": "Genius ideas", "reply_emoji": "🧐", "starling_key": "dm_suggestreply_genius"}
  ],
  "Home & Garden": [
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"},
    {"reply_text": "Beautiful", "reply_emoji": "🌟", "starling_key": "dm_suggestreply_beautiful"},
    {"reply_text": "So interesting", "reply_emoji": "🏡", "starling_key": "dm_suggestreply_so_interesting"},
    {"reply_text": "That's wild", "reply_emoji": "🤠", "starling_key": "dm_suggestreply_thats_wild"},
    {"reply_text": "So cool!", "reply_emoji": "💯", "starling_key": "dm_suggestreply_so_cool3"}
  ],
  "Toys & Collectables": [
    {"reply_text": "Nostalgia feels", "reply_emoji": "🕹️", "starling_key": "dm_suggestreply_nostalgia_feels"},
    {"reply_text": "Whoa!", "reply_emoji": "😮", "starling_key": "dm_suggestreply_whoa"},
    {"reply_text": "I want one", "reply_emoji": "😍", "starling_key": "dm_suggestreply_i_want_one"},
    {"reply_text": "So cool!", "reply_emoji": "👌", "starling_key": "dm_suggestreply_so_cool4"},
    {"reply_text": "Huge fan", "reply_emoji": "😵", "starling_key": "dm_suggestreply_huge_fan3"}
  ],
  "Selfies": [
    {"reply_text": "Great shot!", "reply_emoji": "📸", "starling_key": "dm_suggestreply_great_shot"},
    {"reply_text": "Selfie game strong", "reply_emoji": "👀", "starling_key": "dm_suggestreply_selfie_game_strong"},
    {"reply_text": "HAHA", "reply_emoji": "😂", "starling_key": "dm_suggestreply_haha"},
    {"reply_text": "Whoa!", "reply_emoji": "😮", "starling_key": "dm_suggestreply_whoa"},
    {"reply_text": "Iconic", "reply_emoji": "👑", "starling_key": "dm_suggestreply_iconic"}
  ],
  "Oddly Satisfying": [
    {"reply_text": "So satisfying!", "reply_emoji": "🤤", "starling_key": "dm_suggestreply_so_satisfying"},
    {"reply_text": "This is great", "reply_emoji": "💯", "starling_key": "dm_suggestreply_this_is_great"},
    {"reply_text": "There's no way", "reply_emoji": "😮", "starling_key": "dm_suggestreply_theres_no_way"},
    {"reply_text": "Mesmerizing", "reply_emoji": "🌀", "starling_key": "dm_suggestreply_mesmerizing"},
    {"reply_text": "Mind-boggling...", "reply_emoji": "🤔", "starling_key": "dm_suggestreply_mind_boggling"}
  ],
  "Street Interviews & Social Experiments": [
    {"reply_text": "Cringe", "reply_emoji": "😬", "starling_key": "dm_suggestreply_cringe"},
    {"reply_text": "Wait a sec", "reply_emoji": "👀", "starling_key": "dm_suggestreply_wait_a_sec"},
    {"reply_text": "Intriguing", "reply_emoji": "🤓", "starling_key": "dm_suggestreply_intriguing"},
    {"reply_text": "Speak up", "reply_emoji": "🗣", "starling_key": "dm_suggestreply_speak_up"},
    {"reply_text": "The audacity", "reply_emoji": "😳", "starling_key": "dm_suggestreply_the_audacity"}
  ],
  "Environmental Protection": [
    {"reply_text": "Save the planet!", "reply_emoji": "🌱", "starling_key": "dm_suggestreply_save_the_planet"},
    {"reply_text": "Speak up", "reply_emoji": "🗣", "starling_key": "dm_suggestreply_speak_up"},
    {"reply_text": "Earth day everyday", "reply_emoji": "🌎", "starling_key": "dm_suggestreply_earth_day_everyday"},
    {"reply_text": "Shocking!", "reply_emoji": "😳", "starling_key": "dm_suggestreply_shocking"},
    {"reply_text": "So fascinating", "reply_emoji": "🤓", "starling_key": "dm_suggestreply_so_fascinating"}
  ],
  "Social News": [
    {"reply_text": "Interesting", "reply_emoji": "🤔", "starling_key": "dm_suggestreply_interesting"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"},
    {"reply_text": "Speak up!", "reply_emoji": "🗣", "starling_key": "dm_suggestreply_speak_up2"},
    {"reply_text": "Good to know!", "reply_emoji": "👍", "starling_key": "dm_suggestreply_good_to_know"},
    {"reply_text": "Cringe", "reply_emoji": "😬", "starling_key": "dm_suggestreply_cringe"}
  ],
  "Social Issues": [
    {"reply_text": "Let's discuss!", "reply_emoji": "🗣️", "starling_key": "dm_suggestreply_lets_discuss"},
    {"reply_text": "Cringe", "reply_emoji": "😬", "starling_key": "dm_suggestreply_cringe"},
    {"reply_text": "I'm speechless!", "reply_emoji": "😶", "starling_key": "dm_suggestreply_im_speechless1"},
    {"reply_text": "So insightful!", "reply_emoji": "👌", "starling_key": "dm_suggestreply_so_insightful"},
    {"reply_text": "That makes me so sad", "reply_emoji": "😢", "starling_key": "dm_suggestreply_that_makes_me_so_sad"}
  ],
   "Beauty": [
    {"reply_text": "Gorgeous look", "reply_emoji": "💋", "starling_key": "dm_suggestreply_gorgeous_look"},
    {"reply_text": "I need!", "reply_emoji": "😩", "starling_key": "dm_suggestreply_i_need2"},
    {"reply_text": "I must have it", "reply_emoji": "🛍️", "starling_key": "dm_suggestreply_i_must_have_it"},
    {"reply_text": "Lovely colors", "reply_emoji": "🌸", "starling_key": "dm_suggestreply_lovely_colors"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"}
  ],
  "Hair": [
    {"reply_text": "So cute!", "reply_emoji": "😍", "starling_key": "dm_suggestreply_so_cute1"},
    {"reply_text": "Amazing hairstyle", "reply_emoji": "😮", "starling_key": "dm_suggestreply_amazing_hairstyle"},
    {"reply_text": "So creative", "reply_emoji": "🎨", "starling_key": "dm_suggestreply_so_creative1"},
    {"reply_text": "Love this!", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this2"},
    {"reply_text": "That transformation", "reply_emoji": "💯", "starling_key": "dm_suggestreply_that_transformation"}
  ],
  "Nail Art": [
    {"reply_text": "Cute", "reply_emoji": "😍", "starling_key": "dm_suggestreply_cute"},
    {"reply_text": "So creative", "reply_emoji": "🎨", "starling_key": "dm_suggestreply_so_creative1"},
    {"reply_text": "Saving this!", "reply_emoji": "🤝", "starling_key": "dm_suggestreply_saving_this1"},
    {"reply_text": "Beautiful", "reply_emoji": "🌟", "starling_key": "dm_suggestreply_beautiful"},
    {"reply_text": "Stunning!", "reply_emoji": "💅", "starling_key": "dm_suggestreply_stunning1"}
  ],
  "Other Beauty": [
    {"reply_text": "Gorgeous look", "reply_emoji": "💋", "starling_key": "dm_suggestreply_gorgeous_look"},
    {"reply_text": "I need!", "reply_emoji": "😩", "starling_key": "dm_suggestreply_i_need2"},
    {"reply_text": "I must have it", "reply_emoji": "🛍️", "starling_key": "dm_suggestreply_i_must_have_it"},
    {"reply_text": "Lovely colors", "reply_emoji": "🌸", "starling_key": "dm_suggestreply_lovely_colors"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"}
  ],
  "Outfits": [
    {"reply_text": "I want one", "reply_emoji": "🛍️", "starling_key": "dm_suggestreply_i_want_one1"},
    {"reply_text": "I need!", "reply_emoji": "😩", "starling_key": "dm_suggestreply_i_need2"},
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"},
    {"reply_text": "Whoa!", "reply_emoji": "😮", "starling_key": "dm_suggestreply_whoa"},
    {"reply_text": "Stunning!", "reply_emoji": "😍", "starling_key": "dm_suggestreply_stunning2"}
  ],
  "Other Fashion": [
    {"reply_text": "I need", "reply_emoji": "😩", "starling_key": "dm_suggestreply_i_need4"},
    {"reply_text": "Love this!", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this2"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"},
    {"reply_text": "Whoa!", "reply_emoji": "😮", "starling_key": "dm_suggestreply_whoa"},
    {"reply_text": "Stunning", "reply_emoji": "😍", "starling_key": "dm_suggestreply_stunning3"}
  ],
  "Movies & TV works": [
    {"reply_text": "HAHA", "reply_emoji": "😂", "starling_key": "dm_suggestreply_haha"},
    {"reply_text": "OMG!", "reply_emoji": "😲", "starling_key": "dm_suggestreply_omg2"},
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"},
    {"reply_text": "Too good!", "reply_emoji": "💯", "starling_key": "dm_suggestreply_too_good1"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"}
  ],
  "Music": [
    {"reply_text": "Love this song", "reply_emoji": "🎶", "starling_key": "dm_suggestreply_love_this_song"},
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"},
    {"reply_text": "A certified banger!", "reply_emoji": "💯", "starling_key": "dm_suggestreply_a_certified_banger"},
    {"reply_text": "Iconic!", "reply_emoji": "👑", "starling_key": "dm_suggestreply_iconic1"}
  ],
  "Theatre & Stage": [
    {"reply_text": "So talented!", "reply_emoji": "👏", "starling_key": "dm_suggestreply_so_talented"},
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"},
    {"reply_text": "Whoa!", "reply_emoji": "😮", "starling_key": "dm_suggestreply_whoa"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"},
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"}
  ],
  "Entertainment News": [
    {"reply_text": "HAHA", "reply_emoji": "😂", "starling_key": "dm_suggestreply_haha"},
    {"reply_text": "No way!", "reply_emoji": "😳", "starling_key": "dm_suggestreply_no_way"},
    {"reply_text": "That's wild", "reply_emoji": "🤠", "starling_key": "dm_suggestreply_thats_wild"},
    {"reply_text": "Good to know!", "reply_emoji": "👍", "starling_key": "dm_suggestreply_good_to_know"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"}
  ],
  "Celebrity Clips & Variety Show": [
    {"reply_text": "Big fan", "reply_emoji": "👏", "starling_key": "dm_suggestreply_big_fan"},
    {"reply_text": "OMG!", "reply_emoji": "😲", "starling_key": "dm_suggestreply_omg2"},
    {"reply_text": "Legend", "reply_emoji": "💯", "starling_key": "dm_suggestreply_legend"},
    {"reply_text": "Huge fan!", "reply_emoji": "😵", "starling_key": "dm_suggestreply_huge_fan2"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"}
  ],
  "Cosplay": [
    {"reply_text": "Mood", "reply_emoji": "🥳", "starling_key": "dm_suggestreply_mood"},
    {"reply_text": "Love this!", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this2"},
    {"reply_text": "The outfits", "reply_emoji": "😲", "starling_key": "dm_suggestreply_the_outfits"},
    {"reply_text": "Love the look!", "reply_emoji": "😍", "starling_key": "dm_suggestreply_love_the_look"},
    {"reply_text": "Slay", "reply_emoji": "😊", "starling_key": "dm_suggestreply_slay"}
  ],
  "Comics & Cartoon, Anime": [
    {"reply_text": "Too good!", "reply_emoji": "💯", "starling_key": "dm_suggestreply_too_good1"},
    {"reply_text": "Love the look!", "reply_emoji": "😍", "starling_key": "dm_suggestreply_love_the_look"},
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"},
    {"reply_text": "So creative", "reply_emoji": "😮", "starling_key": "dm_suggestreply_so_creative2"},
    {"reply_text": "LOL", "reply_emoji": "😂", "starling_key": "dm_suggestreply_lol"}
  ],
  "Video Games": [
    {"reply_text": "Love this game", "reply_emoji": "🎮", "starling_key": "dm_suggestreply_love_this_game"},
    {"reply_text": "Looks so cool", "reply_emoji": "💯", "starling_key": "dm_suggestreply_looks_so_cool"},
    {"reply_text": "I want!", "reply_emoji": "😮", "starling_key": "dm_suggestreply_i_want3"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"},
    {"reply_text": "Whoa!", "reply_emoji": "😮", "starling_key": "dm_suggestreply_whoa"}
  ],
  "Non-Video Games": [
    {"reply_text": "So cool", "reply_emoji": "😎", "starling_key": "dm_suggestreply_so_cool2"},
    {"reply_text": "I want to try!", "reply_emoji": "🤩", "starling_key": "dm_suggestreply_i_want_to_try"},
    {"reply_text": "LOL", "reply_emoji": "😂", "starling_key": "dm_suggestreply_lol"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"},
    {"reply_text": "Looks great", "reply_emoji": "💯", "starling_key": "dm_suggestreply_looks_great1"}
  ],
  "Supernatural & Horror": [
    {"reply_text": "I'm shook", "reply_emoji": "😱", "starling_key": "dm_suggestreply_im_shook"},
    {"reply_text": "OMG!", "reply_emoji": "😲", "starling_key": "dm_suggestreply_omg2"},
    {"reply_text": "Are you serious", "reply_emoji": "😭", "starling_key": "dm_suggestreply_are_you_serious1"},
    {"reply_text": "So scary", "reply_emoji": "😱", "starling_key": "dm_suggestreply_so_scary"},
    {"reply_text": "No way!", "reply_emoji": "😳", "starling_key": "dm_suggestreply_no_way"}
  ],
  "Traditional Culture": [
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"},
    {"reply_text": "Interesting", "reply_emoji": "🤔", "starling_key": "dm_suggestreply_interesting"},
    {"reply_text": "Fascinating culture", "reply_emoji": "👏", "starling_key": "dm_suggestreply_fascinating_culture"},
    {"reply_text": "So cool", "reply_emoji": "👍", "starling_key": "dm_suggestreply_so_cool5"}
  ],
  "School Education": [
    {"reply_text": "Learning so much", "reply_emoji": "🧠", "starling_key": "dm_suggestreply_learning_so_much"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"},
    {"reply_text": "Back to school!", "reply_emoji": "📚", "starling_key": "dm_suggestreply_back_to_school"},
    {"reply_text": "So interesting", "reply_emoji": "🧐", "starling_key": "dm_suggestreply_so_interesting2"},
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"}
  ],
  "Professional & Personal Development": [
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"},
    {"reply_text": "So interesting", "reply_emoji": "👍", "starling_key": "dm_suggestreply_so_interesting3"},
    {"reply_text": "Valuable tips", "reply_emoji": "💯", "starling_key": "dm_suggestreply_valuable_tips"},
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"},
    {"reply_text": "Thought-provoking!", "reply_emoji": "🧠", "starling_key": "dm_suggestreply_thought_provoking"}
  ],
  "Humanities": [
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"},
    {"reply_text": "Interesting topic", "reply_emoji": "🧐", "starling_key": "dm_suggestreply_interesting_topic"},
    {"reply_text": "Thought-provoking!", "reply_emoji": "🧠", "starling_key": "dm_suggestreply_thought_provoking"},
    {"reply_text": "So cool", "reply_emoji": "👍", "starling_key": "dm_suggestreply_so_cool5"},
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"}
  ],
  "Business & Finance": [
    {"reply_text": "Great advice!", "reply_emoji": "💼", "starling_key": "dm_suggestreply_great_advice"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"},
    {"reply_text": "Valuable tips", "reply_emoji": "💯", "starling_key": "dm_suggestreply_valuable_tips"},
    {"reply_text": "Interesting", "reply_emoji": "🤔", "starling_key": "dm_suggestreply_interesting"},
    {"reply_text": "Saving this", "reply_emoji": "🤝", "starling_key": "dm_suggestreply_saving_this"}
  ],
  "Science": [
    {"reply_text": "Fascinating", "reply_emoji": "🤓", "starling_key": "dm_suggestreply_fascinating"},
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"},
    {"reply_text": "Amazing technology!", "reply_emoji": "🤖", "starling_key": "dm_suggestreply_amazing_technology"},
    {"reply_text": "OMG!", "reply_emoji": "😲", "starling_key": "dm_suggestreply_omg2"}
  ],
  "Motivation": [
    {"reply_text": "So inspiring", "reply_emoji": "👍", "starling_key": "dm_suggestreply_so_inspiring"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"},
    {"reply_text": "Send me more", "reply_emoji": "🤝", "starling_key": "dm_suggestreply_send_me_more"},
    {"reply_text": "So uplifting!", "reply_emoji": "🤗", "starling_key": "dm_suggestreply_so_uplifting"},
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"}
  ],
  "Tech Products & Infos": [
    {"reply_text": "Cool tech!", "reply_emoji": "📱", "starling_key": "dm_suggestreply_cool_tech"},
    {"reply_text": "Whoa!", "reply_emoji": "😮", "starling_key": "dm_suggestreply_whoa"},
    {"reply_text": "I want one", "reply_emoji": "🤓", "starling_key": "dm_suggestreply_i_want_one2"},
    {"reply_text": "I need!", "reply_emoji": "😩", "starling_key": "dm_suggestreply_i_need2"},
    {"reply_text": "So cool", "reply_emoji": "👍", "starling_key": "dm_suggestreply_so_cool5"}
  ],
  "Software & APPs": [
    {"reply_text": "Whoa!", "reply_emoji": "😮", "starling_key": "dm_suggestreply_whoa"},
    {"reply_text": "I need", "reply_emoji": "😩", "starling_key": "dm_suggestreply_i_need4"},
    {"reply_text": "Cool app!", "reply_emoji": "📱", "starling_key": "dm_suggestreply_cool_app"},
    {"reply_text": "So cool", "reply_emoji": "👍", "starling_key": "dm_suggestreply_so_cool5"},
    {"reply_text": "This is amazing", "reply_emoji": "🤯", "starling_key": "dm_suggestreply_this_is_amazing"}
  ],
  "Photography": [
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"},
    {"reply_text": "Iconic", "reply_emoji": "👑", "starling_key": "dm_suggestreply_iconic"},
    {"reply_text": "Great shots", "reply_emoji": "📷", "starling_key": "dm_suggestreply_great_shots"},
    {"reply_text": "Inspired", "reply_emoji": "🤓", "starling_key": "dm_suggestreply_inspired"},
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"}
  ],
  "Traditional Sports": [
    {"reply_text": "HAHA", "reply_emoji": "😂", "starling_key": "dm_suggestreply_haha"},
    {"reply_text": "Legend", "reply_emoji": "💯", "starling_key": "dm_suggestreply_legend"},
    {"reply_text": "OMG!", "reply_emoji": "😲", "starling_key": "dm_suggestreply_omg2"},
    {"reply_text": "Cool game", "reply_emoji": "👍", "starling_key": "dm_suggestreply_cool_game"},
    {"reply_text": "Too fit", "reply_emoji": "💪", "starling_key": "dm_suggestreply_too_fit"}
  ],
  "Extreme Sports": [
    {"reply_text": "I wish I could do that", "reply_emoji": "🥲", "starling_key": "dm_suggestreply_i_wish_i_could_do_that"},
    {"reply_text": "Too fit", "reply_emoji": "💪", "starling_key": "dm_suggestreply_too_fit"},
    {"reply_text": "I'm shook", "reply_emoji": "😱", "starling_key": "dm_suggestreply_im_shook"},
    {"reply_text": "Mindblown!", "reply_emoji": "🤯", "starling_key": "dm_suggestreply_mindblown"},
    {"reply_text": "No way!", "reply_emoji": "😳", "starling_key": "dm_suggestreply_no_way"}
  ],
  "Sports News": [
    {"reply_text": "Legend", "reply_emoji": "💯", "starling_key": "dm_suggestreply_legend"},
    {"reply_text": "Interesting", "reply_emoji": "🤔", "starling_key": "dm_suggestreply_interesting"},
    {"reply_text": "Let's go", "reply_emoji": "🙌", "starling_key": "dm_suggestreply_lets_go"},
    {"reply_text": "I hope so", "reply_emoji": "🤞", "starling_key": "dm_suggestreply_i_hope_so"},
    {"reply_text": "There's no way", "reply_emoji": "😮", "starling_key": "dm_suggestreply_theres_no_way"}
  ],
  "Fitness": [
    {"reply_text": "Too fit", "reply_emoji": "💪", "starling_key": "dm_suggestreply_too_fit"},
    {"reply_text": "Respect", "reply_emoji": "👊", "starling_key": "dm_suggestreply_respect"},
    {"reply_text": "Whoa!", "reply_emoji": "😮", "starling_key": "dm_suggestreply_whoa"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"},
    {"reply_text": "Inspired", "reply_emoji": "🤓", "starling_key": "dm_suggestreply_inspired"}
  ],
  "Health & Wellness": [
    {"reply_text": "Manifesting this", "reply_emoji": "✨", "starling_key": "dm_suggestreply_manifesting_this"},
    {"reply_text": "Saving this", "reply_emoji": "🤝", "starling_key": "dm_suggestreply_saving_this"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"},
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"},
    {"reply_text": "Need this", "reply_emoji": "🌱", "starling_key": "dm_suggestreply_need_this"}
  ],
  "Fishing, Hunting & Camping": [
    {"reply_text": "That's wild", "reply_emoji": "🤠", "starling_key": "dm_suggestreply_thats_wild"},
    {"reply_text": "I'm down", "reply_emoji": "👍", "starling_key": "dm_suggestreply_im_down"},
    {"reply_text": "So cool", "reply_emoji": "😎", "starling_key": "dm_suggestreply_so_cool2"},
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"},
    {"reply_text": "Let's go", "reply_emoji": "🌲", "starling_key": "dm_suggestreply_lets_go2"}
  ],
  "Cars, Trucks & Motorcycles": [
    {"reply_text": "I want one", "reply_emoji": "👌", "starling_key": "dm_suggestreply_i_want_one3"},
    {"reply_text": "Stunning", "reply_emoji": "😍", "starling_key": "dm_suggestreply_stunning3"},
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"},
    {"reply_text": "So cool!", "reply_emoji": "😎", "starling_key": "dm_suggestreply_so_cool1"},
    {"reply_text": "Whoa!", "reply_emoji": "😮", "starling_key": "dm_suggestreply_whoa"}
  ],
  "Other Transportation": [
    {"reply_text": "Interesting", "reply_emoji": "🤔", "starling_key": "dm_suggestreply_interesting"},
    {"reply_text": "Thanks for sharing", "reply_emoji": "🙏", "starling_key": "dm_suggestreply_thanks_for_sharing"},
    {"reply_text": "Love this", "reply_emoji": "❤️", "starling_key": "dm_suggestreply_love_this"},
    {"reply_text": "So cool", "reply_emoji": "😎", "starling_key":"dm_suggestreply_so_cool2"},
    {"reply_text": "Whoa", "reply_emoji": "😮", "starling_key":"dm_suggestreply_whoa"}
  ]
}
`

var category = `
{
  "10021": "Other_Animal",
  "10057": "Non_VideoGames",
  "10067": "TechProducts_Infos",
  "10014": "Diary_VLOG",
  "10013": "Motivation",
  "10087": "Social_Issues",
  "10068": "School_Education",
  "10040": "Cooking",
  "10058": "Extreme_Sports",
  "10051": "Cosplay",
  "10006": "Magic",
  "10032": "Other_Fashion",
  "10050": "Toys_Ornaments",
  "10056": "VideoGames",
  "10011": "Other_Talent",
  "10034": "Life_Hacks",
  "10075": "Random_Shoot",
  "10012": "Romance",
  "10036": "Other_Art",
  "10039": "Food_Display",
  "10090": "Traditional_Culture",
  "10093": "Business_Finance",
  "10015": "Campus_Life",
  "10000": "Scripted_Comedy",
  "10019": "Pets",
  "10029": "Outfit",
  "10017": "Baby",
  "10060": "Sport_News",
  "10022": "Environment_Protection",
  "10037": "Oddly_Satisfying",
  "10042": "Drinks",
  "10045": "Movies_TVworks",
  "10002": "Hilarious_Fails",
  "10062": "Cars_Trucks_Motorcycles",
  "10092": "Humanities",
  "10003": "Other_Comedy",
  "10089": "Supernatural_Horror",
  "10028": "Nail_Art",
  "10026": "Beauty",
  "10080": "FingerDance_BasicDance",
  "10071": "Lipsync",
  "10025": "Hair",
  "10018": "Family",
  "10027": "Other_Beauty",
  "10004": "Scripted_Drama",
  "10035": "Graphic_Art",
  "10020": "Farm_Animals",
  "10074": "Advertisement",
  "10009": "Singing_Instruments",
  "10052": "Comics_Cartoon_Anime",
  "10066": "Scenery_Plants",
  "10063": "Other_Transportation",
  "10041": "Mukbangs_Tasting",
  "10081": "Interviews_Experiments",
  "10044": "Home_Garden",
  "10046": "Music",
  "10086": "Recreation_Facilities",
  "10049": "Entertainment_News",
  "10070": "Work_Jobs",
  "10094": "Science",
  "10033": "DIY_Handcrafts",
  "10047": "Theatre_Stage",
  "10091": "PersonalDevelopment",
  "10005": "Dance",
  "10061": "Fitness",
  "10082": "Photography",
  "10059": "Traditional_Sports",
  "10073": "Selfie",
  "10008": "Professional_SpecialEffects",
  "10095": "Software_APPs",
  "10083": "No_classification",
  "10085": "FoodTour_Recommendations",
  "10024": "Social_News",
  "10043": "Travel",
  "10088": "CelebrityClips_VarietyShow",
  "10001": "Prank",
  "10096": "Health_Wellness",
  "10064": "Fishing_Hunting_Camping"
}
`

func SugReply() {
	categoryMap := make(map[string]string)
	rawDataMap := make(map[string]interface{})

	jsonx.UnmarshalFromString(category, &categoryMap)
	jsonx.UnmarshalFromString(rawData, &rawDataMap)

	fmt.Println("category length : %v", len(categoryMap))
	fmt.Println("rawdata length : %v", len(rawDataMap))

	for key2, value2 := range categoryMap {
		value2 = strings.ReplaceAll(value2, "_", "")
		categoryMap[key2] = strings.ToLower(value2)
	}

	for key, value := range rawDataMap {
		key3 := strings.ReplaceAll(key, " ", "")
		key3 = strings.ReplaceAll(key3, "&", "")
		key3 = strings.ReplaceAll(key3, ",", "")
		key3 = strings.ReplaceAll(key3, "-", "")
		delete(rawDataMap, key)
		rawDataMap[strings.ToLower(key3)] = value

	}

	for key, value := range rawDataMap {
		for key2, value2 := range categoryMap {
			if value2 == key {
				rawDataMap[key2] = value
				delete(rawDataMap, key)
			}
		}
	}
	fmt.Println("rawdata after length : %v", len(rawDataMap))

	fmt.Println(jsonx.ToString(rawDataMap))
}
