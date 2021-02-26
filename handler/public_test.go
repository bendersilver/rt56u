package handler

import (
	"bytes"
	"net/http/httptest"
	"strings"
	"testing"
)

var rd = strings.NewReader(`#EXTM3U
#EXTINF:-1 tvg-id="303" tvg-name="Nickelodeon HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/303/original/590800.png" group-title="Детские",Nickelodeon HD
http://playlist.tv.planeta.tc/playlist/hls/303-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="590" tvg-name="Капитан Фантастика" tvg-logo="http://gl.weburg.net/00/tv/channels/1/590/original/3862265.png" group-title="Детские",Капитан Фантастика
http://playlist.tv.planeta.tc/playlist/hls/590-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="855" tvg-name="Мульт HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/855/original/7136407.png" group-title="Детские",Мульт HD
http://playlist.tv.planeta.tc/playlist/hls/855-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1253" tvg-name="СТС Kids HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1253/original/7483325.png" group-title="Детские",СТС Kids HD
http://playlist.tv.planeta.tc/playlist/hls/1253-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="128" tvg-name="Boomerang" tvg-logo="http://gl.weburg.net/00/tv/channels/1/128/original/595233.png" group-title="Детские",Boomerang
http://playlist.tv.planeta.tc/playlist/hls/128-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="165" tvg-name="Disney" tvg-logo="http://gl.weburg.net/00/tv/channels/1/165/original/595235.png" group-title="Детские",Disney
http://playlist.tv.planeta.tc/playlist/hls/165-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="409" tvg-name="Nick Jr." tvg-logo="http://gl.weburg.net/00/tv/channels/1/409/original/591504.png" group-title="Детские",Nick Jr.
http://playlist.tv.planeta.tc/playlist/hls/409-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="206" tvg-name="Карусель" tvg-logo="http://gl.weburg.net/00/tv/channels/1/206/original/595239.png" group-title="Детские",Карусель
http://playlist.tv.planeta.tc/playlist/hls/206-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1196" tvg-name="Малыш ТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1196/original/7429728.png" group-title="Детские",Малыш ТВ
http://playlist.tv.planeta.tc/playlist/hls/1196-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="257" tvg-name="Радость Моя" tvg-logo="http://gl.weburg.net/00/tv/channels/1/257/original/595241.png" group-title="Детские",Радость Моя
http://playlist.tv.planeta.tc/playlist/hls/257-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="80" tvg-name="Рыжий" tvg-logo="http://gl.weburg.net/00/tv/channels/1/80/original/595237.png" group-title="Детские",Рыжий
http://playlist.tv.planeta.tc/playlist/hls/80-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="230" tvg-name="Animal Planet HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/230/original/590792.png" group-title="Живая природа",Animal Planet HD
http://playlist.tv.planeta.tc/playlist/hls/230-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="220" tvg-name="Nat Geo Wild HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/220/original/590796.png" group-title="Живая природа",Nat Geo Wild HD
http://playlist.tv.planeta.tc/playlist/hls/220-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="498" tvg-name="В мире животных HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/498/original/939310.png" group-title="Живая природа",В мире животных HD
http://playlist.tv.planeta.tc/playlist/hls/498-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="599" tvg-name="Живая природа HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/599/original/4266190.png" group-title="Живая природа",Живая природа HD
http://playlist.tv.planeta.tc/playlist/hls/599-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1239" tvg-name="Наша Сибирь 4К" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1239/original/7468471.png" group-title="Живая природа",Наша Сибирь 4К
http://playlist.tv.planeta.tc/playlist/hls/1239-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="119" tvg-name="Nat Geo Wild" tvg-logo="http://gl.weburg.net/00/tv/channels/1/119/original/591600.png" group-title="Живая природа",Nat Geo Wild
http://playlist.tv.planeta.tc/playlist/hls/119-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="468" tvg-name="Viasat Nature" tvg-logo="http://gl.weburg.net/00/tv/channels/1/468/original/645240.png" group-title="Живая природа",Viasat Nature
http://playlist.tv.planeta.tc/playlist/hls/468-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="77" tvg-name="Домашние животные" tvg-logo="http://gl.weburg.net/00/tv/channels/1/77/original/591610.png" group-title="Живая природа",Домашние животные
http://playlist.tv.planeta.tc/playlist/hls/77-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1300" tvg-name="365 дней ТВ HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1300/original/7498970.png" group-title="Познавательные",365 дней ТВ HD
http://playlist.tv.planeta.tc/playlist/hls/1300-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="231" tvg-name="Discovery Channel HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/231/original/590788.png" group-title="Познавательные",Discovery Channel HD
http://playlist.tv.planeta.tc/playlist/hls/231-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="615" tvg-name="English Club TV HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/615/original/4953486.png" group-title="Познавательные",English Club TV HD
http://playlist.tv.planeta.tc/playlist/hls/615-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="182" tvg-name="HDL" tvg-logo="http://gl.weburg.net/00/tv/channels/1/182/original/590492.png" group-title="Познавательные",HDL
http://playlist.tv.planeta.tc/playlist/hls/182-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1212" tvg-name="HISTORY2 HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1212/original/7442774.png" group-title="Познавательные",HISTORY2 HD
http://playlist.tv.planeta.tc/playlist/hls/1212-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1230" tvg-name="ID Investigation Discovery HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1230/original/7461744.png" group-title="Познавательные",ID Investigation Discovery HD
http://playlist.tv.planeta.tc/playlist/hls/1230-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1218" tvg-name="Insight UHD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1218/original/7453474.png" group-title="Познавательные",Insight UHD
http://playlist.tv.planeta.tc/playlist/hls/1218-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1312" tvg-name="LOVE NATURE 4K" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1312/original/7502718.png" group-title="Познавательные",LOVE NATURE 4K
http://playlist.tv.planeta.tc/playlist/hls/1312-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1321" tvg-name="MUSEUM 4K" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1321/original/7504296.png" group-title="Познавательные",MUSEUM 4K
http://playlist.tv.planeta.tc/playlist/hls/1321-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1313" tvg-name="MUSEUM HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1313/original/7502723.png" group-title="Познавательные",MUSEUM HD
http://playlist.tv.planeta.tc/playlist/hls/1313-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="219" tvg-name="National Geographic HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/219/original/590794.png" group-title="Познавательные",National Geographic HD
http://playlist.tv.planeta.tc/playlist/hls/219-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="438" tvg-name="RTД HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/438/original/603466.png" group-title="Познавательные",RTД HD
http://playlist.tv.planeta.tc/playlist/hls/438-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="492" tvg-name="RTД HD на русском" tvg-logo="http://gl.weburg.net/00/tv/channels/1/492/original/893574.png" group-title="Познавательные",RTД HD на русском
http://playlist.tv.planeta.tc/playlist/hls/492-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="450" tvg-name="TLC HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/450/original/621301.png" group-title="Познавательные",TLC HD
http://playlist.tv.planeta.tc/playlist/hls/450-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="442" tvg-name="Travel Channel HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/442/original/614129.png" group-title="Познавательные",Travel Channel HD
http://playlist.tv.planeta.tc/playlist/hls/442-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1273" tvg-name="Viasat Explore HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1273/original/7494719.png" group-title="Познавательные",Viasat Explore HD
http://playlist.tv.planeta.tc/playlist/hls/1273-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="533" tvg-name="Viasat History HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/533/original/2053009.png" group-title="Познавательные",Viasat History HD
http://playlist.tv.planeta.tc/playlist/hls/533-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1274" tvg-name="Viasat Nature HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1274/original/7494722.png" group-title="Познавательные",Viasat Nature HD
http://playlist.tv.planeta.tc/playlist/hls/1274-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1318" tvg-name="Арсенал HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1318/original/7503390.png" group-title="Познавательные",Арсенал HD
http://playlist.tv.planeta.tc/playlist/hls/1318-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="477" tvg-name="Моя планета HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/477/original/660884.png" group-title="Познавательные",Моя планета HD
http://playlist.tv.planeta.tc/playlist/hls/477-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1299" tvg-name="Наука HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1299/original/7498443.png" group-title="Познавательные",Наука HD
http://playlist.tv.planeta.tc/playlist/hls/1299-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1233" tvg-name="Наша Сибирь HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1233/original/7463145.png" group-title="Познавательные",Наша Сибирь HD
http://playlist.tv.planeta.tc/playlist/hls/1233-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="497" tvg-name="Эврика HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/497/original/939194.png" group-title="Познавательные",Эврика HD
http://playlist.tv.planeta.tc/playlist/hls/497-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="122" tvg-name="365 Дней ТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/122/original/591494.png" group-title="Познавательные",365 Дней ТВ
http://playlist.tv.planeta.tc/playlist/hls/122-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="469" tvg-name="Da Vinci Learning" tvg-logo="http://gl.weburg.net/00/tv/channels/1/469/original/645258.png" group-title="Познавательные",Da Vinci Learning
http://playlist.tv.planeta.tc/playlist/hls/469-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="214" tvg-name="Discovery Science" tvg-logo="http://gl.weburg.net/00/tv/channels/1/214/original/595856.png" group-title="Познавательные",Discovery Science
http://playlist.tv.planeta.tc/playlist/hls/214-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="213" tvg-name="DTX" tvg-logo="http://gl.weburg.net/00/tv/channels/1/213/original/595858.png" group-title="Познавательные",DTX
http://playlist.tv.planeta.tc/playlist/hls/213-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="583" tvg-name="Galaxy" tvg-logo="http://gl.weburg.net/00/tv/channels/1/583/original/3408066.png" group-title="Познавательные",Galaxy
http://playlist.tv.planeta.tc/playlist/hls/583-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="212" tvg-name="ID Investigation Discovery" tvg-logo="http://gl.weburg.net/00/tv/channels/1/212/original/595860.png" group-title="Познавательные",ID Investigation Discovery
http://playlist.tv.planeta.tc/playlist/hls/212-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="308" tvg-name="NANO" tvg-logo="http://gl.weburg.net/00/tv/channels/1/308/original/591645.png" group-title="Познавательные",NANO
http://playlist.tv.planeta.tc/playlist/hls/308-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="36" tvg-name="National Geographic" tvg-logo="http://gl.weburg.net/00/tv/channels/1/36/original/591598.png" group-title="Познавательные",National Geographic
http://playlist.tv.planeta.tc/playlist/hls/36-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="85" tvg-name="Ocean TV" tvg-logo="http://gl.weburg.net/00/tv/channels/1/85/original/591602.png" group-title="Познавательные",Ocean TV
http://playlist.tv.planeta.tc/playlist/hls/85-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="40" tvg-name="Viasat Explore" tvg-logo="http://gl.weburg.net/00/tv/channels/1/40/original/595897.png" group-title="Познавательные",Viasat Explore
http://playlist.tv.planeta.tc/playlist/hls/40-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="41" tvg-name="Viasat History" tvg-logo="http://gl.weburg.net/00/tv/channels/1/41/original/595899.png" group-title="Познавательные",Viasat History
http://playlist.tv.planeta.tc/playlist/hls/41-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="584" tvg-name="ЕГЭ ТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/584/original/3408069.png" group-title="Познавательные",ЕГЭ ТВ
http://playlist.tv.planeta.tc/playlist/hls/584-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="843" tvg-name="Калейдоскоп ТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/843/original/6786083.png" group-title="Познавательные",Калейдоскоп ТВ
http://playlist.tv.planeta.tc/playlist/hls/843-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="832" tvg-name="Министерство идей" tvg-logo="http://gl.weburg.net/00/tv/channels/1/832/original/6363238.png" group-title="Познавательные",Министерство идей
http://playlist.tv.planeta.tc/playlist/hls/832-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="76" tvg-name="Психология 21" tvg-logo="http://gl.weburg.net/00/tv/channels/1/76/original/591661.png" group-title="Познавательные",Психология 21
http://playlist.tv.planeta.tc/playlist/hls/76-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="59" tvg-name="Совершенно секретно" tvg-logo="http://gl.weburg.net/00/tv/channels/1/59/original/591576.png" group-title="Познавательные",Совершенно секретно
http://playlist.tv.planeta.tc/playlist/hls/59-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="447" tvg-name="RTG HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/447/original/618337.png" group-title="Путешествия и туризм",RTG HD
http://playlist.tv.planeta.tc/playlist/hls/447-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="309" tvg-name="Travel+Adventure HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/309/original/590803.png" group-title="Путешествия и туризм",Travel+Adventure HD
http://playlist.tv.planeta.tc/playlist/hls/309-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1241" tvg-name="Большая Азия HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1241/original/7473352.png" group-title="Путешествия и туризм",Большая Азия HD
http://playlist.tv.planeta.tc/playlist/hls/1241-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="158" tvg-name="Приключения HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/158/original/590786.png" group-title="Путешествия и туризм",Приключения HD
http://playlist.tv.planeta.tc/playlist/hls/158-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="120" tvg-name="RTG TV" tvg-logo="http://gl.weburg.net/00/tv/channels/1/120/original/591604.png" group-title="Путешествия и туризм",RTG TV
http://playlist.tv.planeta.tc/playlist/hls/120-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="318" tvg-name="Travel+Adventure" tvg-logo="http://gl.weburg.net/00/tv/channels/1/318/original/591508.png" group-title="Путешествия и туризм",Travel+Adventure
http://playlist.tv.planeta.tc/playlist/hls/318-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1203" tvg-name="Поехали!" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1203/original/7438647.png" group-title="Путешествия и туризм",Поехали!
http://playlist.tv.planeta.tc/playlist/hls/1203-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="79" tvg-name="Телепутешествия" tvg-logo="http://gl.weburg.net/00/tv/channels/1/79/original/591616.png" group-title="Путешествия и туризм",Телепутешествия
http://playlist.tv.planeta.tc/playlist/hls/79-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1231" tvg-name="Discovery Science HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1231/original/7461741.png" group-title="Развлекательные",Discovery Science HD
http://playlist.tv.planeta.tc/playlist/hls/1231-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1246" tvg-name="FAN HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1246/original/7479711.png" group-title="Развлекательные",FAN HD
http://playlist.tv.planeta.tc/playlist/hls/1246-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="448" tvg-name="History HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/448/original/620283.png" group-title="Развлекательные",History HD
http://playlist.tv.planeta.tc/playlist/hls/448-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1217" tvg-name="HOME 4K" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1217/original/7453471.png" group-title="Развлекательные",HOME 4K
http://playlist.tv.planeta.tc/playlist/hls/1217-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1316" tvg-name="Luxe TV HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1316/original/7502734.png" group-title="Развлекательные",Luxe TV HD
http://playlist.tv.planeta.tc/playlist/hls/1316-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1266" tvg-name="Глазами туриста HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1266/original/7493341.png" group-title="Развлекательные",Глазами туриста HD
http://playlist.tv.planeta.tc/playlist/hls/1266-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1252" tvg-name="Телекон 24 HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1252/original/7482172.png" group-title="Развлекательные",Телекон 24 HD
http://playlist.tv.planeta.tc/playlist/hls/1252-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="24" tvg-name="2x2" tvg-logo="http://gl.weburg.net/00/tv/channels/1/24/original/593842.png" group-title="Развлекательные",2x2
http://playlist.tv.planeta.tc/playlist/hls/24-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1247" tvg-name="MCM TOP" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1247/original/7480133.png" group-title="Развлекательные",MCM TOP
http://playlist.tv.planeta.tc/playlist/hls/1247-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1223" tvg-name="TVMChannel" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1223/original/7459391.png" group-title="Развлекательные",TVMChannel
http://playlist.tv.planeta.tc/playlist/hls/1223-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="78" tvg-name="Вопросы и ответы" tvg-logo="http://gl.weburg.net/00/tv/channels/1/78/original/591647.png" group-title="Развлекательные",Вопросы и ответы
http://playlist.tv.planeta.tc/playlist/hls/78-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="46" tvg-name="Время" tvg-logo="http://gl.weburg.net/00/tv/channels/1/46/original/591471.png" group-title="Развлекательные",Время
http://playlist.tv.planeta.tc/playlist/hls/46-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1185" tvg-name="КВН ТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1185/original/7419600.png" group-title="Развлекательные",КВН ТВ
http://playlist.tv.planeta.tc/playlist/hls/1185-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="121" tvg-name="Ностальгия" tvg-logo="http://gl.weburg.net/00/tv/channels/1/121/original/591612.png" group-title="Развлекательные",Ностальгия
http://playlist.tv.planeta.tc/playlist/hls/121-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="440" tvg-name="ПЯТНИЦА" tvg-logo="http://gl.weburg.net/00/tv/channels/1/440/original/604131.png" group-title="Развлекательные",ПЯТНИЦА
http://playlist.tv.planeta.tc/playlist/hls/440-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="75" tvg-name="Ретро" tvg-logo="http://gl.weburg.net/00/tv/channels/1/75/original/591663.png" group-title="Развлекательные",Ретро
http://playlist.tv.planeta.tc/playlist/hls/75-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="613" tvg-name="СТС love" tvg-logo="http://gl.weburg.net/00/tv/channels/1/613/original/4815283.png" group-title="Развлекательные",СТС love
http://playlist.tv.planeta.tc/playlist/hls/613-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1221" tvg-name="Суббота!" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1221/original/7456889.png" group-title="Развлекательные",Суббота!
http://playlist.tv.planeta.tc/playlist/hls/1221-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="20" tvg-name="ТВ3" tvg-logo="http://gl.weburg.net/00/tv/channels/1/20/original/591126.png" group-title="Развлекательные",ТВ3
http://playlist.tv.planeta.tc/playlist/hls/20-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="99" tvg-name="ТДК" tvg-logo="http://gl.weburg.net/00/tv/channels/1/99/original/593851.png" group-title="Развлекательные",ТДК
http://playlist.tv.planeta.tc/playlist/hls/99-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="582" tvg-name="Театр" tvg-logo="http://gl.weburg.net/00/tv/channels/1/582/original/3408063.png" group-title="Развлекательные",Театр
http://playlist.tv.planeta.tc/playlist/hls/582-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="367" tvg-name="ТНТ4" tvg-logo="http://gl.weburg.net/00/tv/channels/1/367/original/591526.png" group-title="Развлекательные",ТНТ4
http://playlist.tv.planeta.tc/playlist/hls/367-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="10" tvg-name="Че" tvg-logo="http://gl.weburg.net/00/tv/channels/1/10/original/591102.png" group-title="Развлекательные",Че
http://playlist.tv.planeta.tc/playlist/hls/10-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="234" tvg-name="Ю-ТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/234/original/593855.png" group-title="Развлекательные",Ю-ТВ
http://playlist.tv.planeta.tc/playlist/hls/234-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="488" tvg-name="1 HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/488/original/841666.png" group-title="Музыка",1 HD
http://playlist.tv.planeta.tc/playlist/hls/488-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1264" tvg-name="AIVA" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1264/original/7492927.png" group-title="Музыка",AIVA
http://playlist.tv.planeta.tc/playlist/hls/1264-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1193" tvg-name="BRIDGE TV DELUXE" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1193/original/7429171.png" group-title="Музыка",BRIDGE TV DELUXE
http://playlist.tv.planeta.tc/playlist/hls/1193-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="625" tvg-name="C MUSIC TV HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/625/original/5448689.png" group-title="Музыка",C MUSIC TV HD
http://playlist.tv.planeta.tc/playlist/hls/625-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1242" tvg-name="Europa Plus TV HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1242/original/7474014.png" group-title="Музыка",Europa Plus TV HD
http://playlist.tv.planeta.tc/playlist/hls/1242-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="218" tvg-name="Mezzo Live HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/218/original/590482.png" group-title="Музыка",Mezzo Live HD
http://playlist.tv.planeta.tc/playlist/hls/218-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="304" tvg-name="MTV Live HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/304/original/590773.png" group-title="Музыка",MTV Live HD
http://playlist.tv.planeta.tc/playlist/hls/304-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1184" tvg-name="Russian MusicBox HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1184/original/7418987.png" group-title="Музыка",Russian MusicBox HD
http://playlist.tv.planeta.tc/playlist/hls/1184-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1315" tvg-name="Stingray iConcerts HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1315/original/7502741.png" group-title="Музыка",Stingray iConcerts HD
http://playlist.tv.planeta.tc/playlist/hls/1315-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="267" tvg-name="Trace Urban HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/267/original/600552.png" group-title="Музыка",Trace Urban HD
http://playlist.tv.planeta.tc/playlist/hls/267-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="28" tvg-name="Bridge TV" tvg-logo="http://gl.weburg.net/00/tv/channels/1/28/original/594715.png" group-title="Музыка",Bridge TV
http://playlist.tv.planeta.tc/playlist/hls/28-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="622" tvg-name="BRIDGE TV CLASSIC" tvg-logo="http://gl.weburg.net/00/tv/channels/1/622/original/5448706.png" group-title="Музыка",BRIDGE TV CLASSIC
http://playlist.tv.planeta.tc/playlist/hls/622-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="434" tvg-name="BRIDGE TV HITS" tvg-logo="http://gl.weburg.net/00/tv/channels/1/434/original/599042.png" group-title="Музыка",BRIDGE TV HITS
http://playlist.tv.planeta.tc/playlist/hls/434-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="207" tvg-name="BRIDGE TV Русский Хит" tvg-logo="http://gl.weburg.net/00/tv/channels/1/207/original/594739.png" group-title="Музыка",BRIDGE TV Русский Хит
http://playlist.tv.planeta.tc/playlist/hls/207-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="365" tvg-name="Club MTV" tvg-logo="http://gl.weburg.net/00/tv/channels/1/365/original/590771.png" group-title="Музыка",Club MTV
http://playlist.tv.planeta.tc/playlist/hls/365-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="229" tvg-name="Europa Plus TV" tvg-logo="http://gl.weburg.net/00/tv/channels/1/229/original/594727.png" group-title="Музыка",Europa Plus TV
http://playlist.tv.planeta.tc/playlist/hls/229-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="363" tvg-name="MTV 80s" tvg-logo="http://gl.weburg.net/00/tv/channels/1/363/original/595907.png" group-title="Музыка",MTV 80s
http://playlist.tv.planeta.tc/playlist/hls/363-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="364" tvg-name="MTV 90s" tvg-logo="http://gl.weburg.net/00/tv/channels/1/364/original/590767.png" group-title="Музыка",MTV 90s
http://playlist.tv.planeta.tc/playlist/hls/364-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="366" tvg-name="MTV HITS" tvg-logo="http://gl.weburg.net/00/tv/channels/1/366/original/590769.png" group-title="Музыка",MTV HITS
http://playlist.tv.planeta.tc/playlist/hls/366-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="357" tvg-name="MTV Россия" tvg-logo="http://gl.weburg.net/00/tv/channels/1/357/original/594731.png" group-title="Музыка",MTV Россия
http://playlist.tv.planeta.tc/playlist/hls/357-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="34" tvg-name="Music Box Ru" tvg-logo="http://gl.weburg.net/00/tv/channels/1/34/original/594733.png" group-title="Музыка",Music Box Ru
http://playlist.tv.planeta.tc/playlist/hls/34-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="37" tvg-name="RU.TV" tvg-logo="http://gl.weburg.net/00/tv/channels/1/37/original/594737.png" group-title="Музыка",RU.TV
http://playlist.tv.planeta.tc/playlist/hls/37-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="302" tvg-name="VH1" tvg-logo="http://gl.weburg.net/00/tv/channels/1/302/original/595904.png" group-title="Музыка",VH1
http://playlist.tv.planeta.tc/playlist/hls/302-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="54" tvg-name="Музыка Первого" tvg-logo="http://gl.weburg.net/00/tv/channels/1/54/original/594751.png" group-title="Музыка",Музыка Первого
http://playlist.tv.planeta.tc/playlist/hls/54-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="27" tvg-name="ТНТ MUSIC" tvg-logo="http://gl.weburg.net/00/tv/channels/1/27/original/594711.png" group-title="Музыка",ТНТ MUSIC
http://playlist.tv.planeta.tc/playlist/hls/27-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="429" tvg-name="Fashion & LifeStyle HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/429/original/590488.png" group-title="Хобби и увлечения",Fashion & LifeStyle HD
http://playlist.tv.planeta.tc/playlist/hls/429-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1220" tvg-name="Fashion & Style 4K" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1220/original/7456730.png" group-title="Хобби и увлечения",Fashion & Style 4K
http://playlist.tv.planeta.tc/playlist/hls/1220-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1263" tvg-name="FashionTV 4K" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1263/original/7491954.png" group-title="Хобби и увлечения",FashionTV 4K
http://playlist.tv.planeta.tc/playlist/hls/1263-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1262" tvg-name="FashionTV HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1262/original/7491950.png" group-title="Хобби и увлечения",FashionTV HD
http://playlist.tv.planeta.tc/playlist/hls/1262-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1206" tvg-name="Food Network HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1206/original/7439906.png" group-title="Хобби и увлечения",Food Network HD
http://playlist.tv.planeta.tc/playlist/hls/1206-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1337" tvg-name="HGTV Home&Garden HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1337/original/7506105.png" group-title="Хобби и увлечения",HGTV Home&Garden HD
http://playlist.tv.planeta.tc/playlist/hls/1337-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="153" tvg-name="MyZen HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/153/original/590765.png" group-title="Хобби и увлечения",MyZen HD
http://playlist.tv.planeta.tc/playlist/hls/153-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1260" tvg-name="MyZen TV 4K" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1260/original/7489913.png" group-title="Хобби и увлечения",MyZen TV 4K
http://playlist.tv.planeta.tc/playlist/hls/1260-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="159" tvg-name="Еда Премиум HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/159/original/590798.png" group-title="Хобби и увлечения",Еда Премиум HD
http://playlist.tv.planeta.tc/playlist/hls/159-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1214" tvg-name="Конный мир HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1214/original/7445842.png" group-title="Хобби и увлечения",Конный мир HD
http://playlist.tv.planeta.tc/playlist/hls/1214-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="421" tvg-name="Охотник и рыболов HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/421/original/593820.png" group-title="Хобби и увлечения",Охотник и рыболов HD
http://playlist.tv.planeta.tc/playlist/hls/421-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="839" tvg-name="E TV" tvg-logo="http://gl.weburg.net/00/tv/channels/1/839/original/6644550.png" group-title="Хобби и увлечения",E TV
http://playlist.tv.planeta.tc/playlist/hls/839-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="8" tvg-name="World Fashion Channel" tvg-logo="http://gl.weburg.net/00/tv/channels/1/8/original/593847.png" group-title="Хобби и увлечения",World Fashion Channel
http://playlist.tv.planeta.tc/playlist/hls/8-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="44" tvg-name="Авто Плюс" tvg-logo="http://gl.weburg.net/00/tv/channels/1/44/original/595867.png" group-title="Хобби и увлечения",Авто Плюс
http://playlist.tv.planeta.tc/playlist/hls/44-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="592" tvg-name="БОБЁР" tvg-logo="http://gl.weburg.net/00/tv/channels/1/592/original/3929097.png" group-title="Хобби и увлечения",БОБЁР
http://playlist.tv.planeta.tc/playlist/hls/592-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="74" tvg-name="Драйв" tvg-logo="http://gl.weburg.net/00/tv/channels/1/74/original/591649.png" group-title="Хобби и увлечения",Драйв
http://playlist.tv.planeta.tc/playlist/hls/74-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="369" tvg-name="Загородная жизнь" tvg-logo="http://gl.weburg.net/00/tv/channels/1/369/original/591651.png" group-title="Хобби и увлечения",Загородная жизнь
http://playlist.tv.planeta.tc/playlist/hls/369-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="50" tvg-name="Здоровое ТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/50/original/591653.png" group-title="Хобби и увлечения",Здоровое ТВ
http://playlist.tv.planeta.tc/playlist/hls/50-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="585" tvg-name="КтоКуда" tvg-logo="http://gl.weburg.net/00/tv/channels/1/585/original/3408072.png" group-title="Хобби и увлечения",КтоКуда
http://playlist.tv.planeta.tc/playlist/hls/585-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="516" tvg-name="Морской" tvg-logo="http://gl.weburg.net/00/tv/channels/1/516/original/1248935.png" group-title="Хобби и увлечения",Морской
http://playlist.tv.planeta.tc/playlist/hls/516-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="126" tvg-name="Мужской" tvg-logo="http://gl.weburg.net/00/tv/channels/1/126/original/591512.png" group-title="Хобби и увлечения",Мужской
http://playlist.tv.planeta.tc/playlist/hls/126-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="56" tvg-name="Охота и рыбалка" tvg-logo="http://gl.weburg.net/00/tv/channels/1/56/original/591657.png" group-title="Хобби и увлечения",Охота и рыбалка
http://playlist.tv.planeta.tc/playlist/hls/56-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1232" tvg-name="Первый вегетарианский" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1232/original/7462917.png" group-title="Хобби и увлечения",Первый вегетарианский
http://playlist.tv.planeta.tc/playlist/hls/1232-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="73" tvg-name="Телекафе" tvg-logo="http://gl.weburg.net/00/tv/channels/1/73/original/593853.png" group-title="Хобби и увлечения",Телекафе
http://playlist.tv.planeta.tc/playlist/hls/73-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="82" tvg-name="Тонус ТВ SD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/82/original/591518.png" group-title="Хобби и увлечения",Тонус ТВ SD
http://playlist.tv.planeta.tc/playlist/hls/82-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="61" tvg-name="Усадьба ТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/61/original/591665.png" group-title="Хобби и увлечения",Усадьба ТВ
http://playlist.tv.planeta.tc/playlist/hls/61-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="125" tvg-name="Успех" tvg-logo="http://gl.weburg.net/00/tv/channels/1/125/original/591516.png" group-title="Хобби и увлечения",Успех
http://playlist.tv.planeta.tc/playlist/hls/125-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1310" tvg-name="Home Shopping Russia" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1310/original/7501677.png" group-title="Телемагазины",Home Shopping Russia
http://playlist.tv.planeta.tc/playlist/hls/1310-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="233" tvg-name="LEOMAX+" tvg-logo="http://gl.weburg.net/00/tv/channels/1/233/original/595839.png" group-title="Телемагазины",LEOMAX+
http://playlist.tv.planeta.tc/playlist/hls/233-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="356" tvg-name="LEOMAX24" tvg-logo="http://gl.weburg.net/00/tv/channels/1/356/original/7347842.png" group-title="Телемагазины",LEOMAX24
http://playlist.tv.planeta.tc/playlist/hls/356-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1201" tvg-name="Shop and Show" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1201/original/7434105.png" group-title="Телемагазины",Shop and Show
http://playlist.tv.planeta.tc/playlist/hls/1201-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1265" tvg-name="Shopping Live" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1265/original/7493028.png" group-title="Телемагазины",Shopping Live
http://playlist.tv.planeta.tc/playlist/hls/1265-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1210" tvg-name="VitrinaTV" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1210/original/7441972.png" group-title="Телемагазины",VitrinaTV
http://playlist.tv.planeta.tc/playlist/hls/1210-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1269" tvg-name="Ювелирочка" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1269/original/7494266.png" group-title="Телемагазины",Ювелирочка
http://playlist.tv.planeta.tc/playlist/hls/1269-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="573" tvg-name="Bollywood HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/573/original/3011549.png" group-title="Мировое кино",Bollywood HD
http://playlist.tv.planeta.tc/playlist/hls/573-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="407" tvg-name="Hollywood HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/407/original/590987.png" group-title="Мировое кино",Hollywood HD
http://playlist.tv.planeta.tc/playlist/hls/407-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1243" tvg-name="Paramount Channel HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1243/original/7476772.png" group-title="Мировое кино",Paramount Channel HD
http://playlist.tv.planeta.tc/playlist/hls/1243-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="830" tvg-name="КИНО ТВ HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/830/original/6363231.png" group-title="Мировое кино",КИНО ТВ HD
http://playlist.tv.planeta.tc/playlist/hls/830-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="38" tvg-name="TV1000" tvg-logo="http://gl.weburg.net/00/tv/channels/1/38/original/595891.png" group-title="Мировое кино",TV1000
http://playlist.tv.planeta.tc/playlist/hls/38-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="148" tvg-name="TV1000 Action" tvg-logo="http://gl.weburg.net/00/tv/channels/1/148/original/595895.png" group-title="Мировое кино",TV1000 Action
http://playlist.tv.planeta.tc/playlist/hls/148-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="141" tvg-name="КИНОМИКС" tvg-logo="http://gl.weburg.net/00/tv/channels/1/141/original/596228.png" group-title="Мировое кино",КИНОМИКС
http://playlist.tv.planeta.tc/playlist/hls/141-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="81" tvg-name="Кинопоказ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/81/original/593833.png" group-title="Мировое кино",Кинопоказ
http://playlist.tv.planeta.tc/playlist/hls/81-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1342" tvg-name="Мосфильм. Золотая коллекция HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1342/original/7507925.png" group-title="Отечественное кино",Мосфильм. Золотая коллекция HD
http://playlist.tv.planeta.tc/playlist/hls/1342-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1305" tvg-name="НАШЕ НОВОЕ КИНО HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1305/original/7498990.png" group-title="Отечественное кино",НАШЕ НОВОЕ КИНО HD
http://playlist.tv.planeta.tc/playlist/hls/1305-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="848" tvg-name="Русский роман HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/848/original/6857702.png" group-title="Отечественное кино",Русский роман HD
http://playlist.tv.planeta.tc/playlist/hls/848-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="39" tvg-name="TV1000 Русское Кино" tvg-logo="http://gl.weburg.net/00/tv/channels/1/39/original/595893.png" group-title="Отечественное кино",TV1000 Русское Кино
http://playlist.tv.planeta.tc/playlist/hls/39-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="49" tvg-name="Дом кино" tvg-logo="http://gl.weburg.net/00/tv/channels/1/49/original/593831.png" group-title="Отечественное кино",Дом кино
http://playlist.tv.planeta.tc/playlist/hls/49-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="227" tvg-name="Любимое кино" tvg-logo="http://gl.weburg.net/00/tv/channels/1/227/original/593835.png" group-title="Отечественное кино",Любимое кино
http://playlist.tv.planeta.tc/playlist/hls/227-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1339" tvg-name="Мосфильм. Золотая коллекция" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1339/original/7506348.png" group-title="Отечественное кино",Мосфильм. Золотая коллекция
http://playlist.tv.planeta.tc/playlist/hls/1339-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1255" tvg-name="Победа" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1255/original/7485546.png" group-title="Отечественное кино",Победа
http://playlist.tv.planeta.tc/playlist/hls/1255-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1345" tvg-name="A1 HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1345/original/7508131.png" group-title="Кино и Сериалы",A1 HD
http://playlist.tv.planeta.tc/playlist/hls/1345-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1346" tvg-name="A2 HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1346/original/7508127.png" group-title="Кино и Сериалы",A2 HD
http://playlist.tv.planeta.tc/playlist/hls/1346-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="360" tvg-name="FOX HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/360/original/590989.png" group-title="Кино и Сериалы",FOX HD
http://playlist.tv.planeta.tc/playlist/hls/360-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="361" tvg-name="Fox Life HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/361/original/590991.png" group-title="Кино и Сериалы",Fox Life HD
http://playlist.tv.planeta.tc/playlist/hls/361-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1258" tvg-name="Paramount Comedy HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1258/original/7488691.png" group-title="Кино и Сериалы",Paramount Comedy HD
http://playlist.tv.planeta.tc/playlist/hls/1258-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="359" tvg-name="Sony Channel HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/359/original/590985.png" group-title="Кино и Сериалы",Sony Channel HD
http://playlist.tv.planeta.tc/playlist/hls/359-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1229" tvg-name="Spike HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1229/original/7461771.png" group-title="Кино и Сериалы",Spike HD
http://playlist.tv.planeta.tc/playlist/hls/1229-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1270" tvg-name="TV1000 Action HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1270/original/7494733.png" group-title="Кино и Сериалы",TV1000 Action HD
http://playlist.tv.planeta.tc/playlist/hls/1270-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1271" tvg-name="TV1000 HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1271/original/7494744.png" group-title="Кино и Сериалы",TV1000 HD
http://playlist.tv.planeta.tc/playlist/hls/1271-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1272" tvg-name="TV1000 Русское Кино HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1272/original/7494741.png" group-title="Кино и Сериалы",TV1000 Русское Кино HD
http://playlist.tv.planeta.tc/playlist/hls/1272-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1222" tvg-name="Ultra HD Cinema" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1222/original/7458390.png" group-title="Кино и Сериалы",Ultra HD Cinema
http://playlist.tv.planeta.tc/playlist/hls/1222-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="586" tvg-name="Дом Кино Премиум HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/586/original/3408075.png" group-title="Кино и Сериалы",Дом Кино Премиум HD
http://playlist.tv.planeta.tc/playlist/hls/586-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1304" tvg-name="КИНОКОМЕДИЯ HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1304/original/7498984.png" group-title="Кино и Сериалы",КИНОКОМЕДИЯ HD
http://playlist.tv.planeta.tc/playlist/hls/1304-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1306" tvg-name="КИНОСЕРИЯ HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1306/original/7498987.png" group-title="Кино и Сериалы",КИНОСЕРИЯ HD
http://playlist.tv.planeta.tc/playlist/hls/1306-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1323" tvg-name="КИНОУЖАС HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1323/original/7505458.png" group-title="Кино и Сериалы",КИНОУЖАС HD
http://playlist.tv.planeta.tc/playlist/hls/1323-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="179" tvg-name="FOX" tvg-logo="http://gl.weburg.net/00/tv/channels/1/179/original/591500.png" group-title="Кино и Сериалы",FOX
http://playlist.tv.planeta.tc/playlist/hls/179-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="178" tvg-name="Fox Life" tvg-logo="http://gl.weburg.net/00/tv/channels/1/178/original/591502.png" group-title="Кино и Сериалы",Fox Life
http://playlist.tv.planeta.tc/playlist/hls/178-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="142" tvg-name="Sony Channel" tvg-logo="http://gl.weburg.net/00/tv/channels/1/142/original/595927.png" group-title="Кино и Сериалы",Sony Channel
http://playlist.tv.planeta.tc/playlist/hls/142-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="109" tvg-name="SONY SCI-FI" tvg-logo="http://gl.weburg.net/00/tv/channels/1/109/original/595923.png" group-title="Кино и Сериалы",SONY SCI-FI
http://playlist.tv.planeta.tc/playlist/hls/109-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="425" tvg-name="SONY TURBO" tvg-logo="http://gl.weburg.net/00/tv/channels/1/425/original/595925.png" group-title="Кино и Сериалы",SONY TURBO
http://playlist.tv.planeta.tc/playlist/hls/425-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1341" tvg-name="НТВ-ХИТ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1341/original/7507477.png" group-title="Кино и Сериалы",НТВ-ХИТ
http://playlist.tv.planeta.tc/playlist/hls/1341-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="163" tvg-name="Eurosport 1 HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/163/original/590977.png" group-title="Спорт",Eurosport 1 HD
http://playlist.tv.planeta.tc/playlist/hls/163-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="427" tvg-name="Eurosport 2 HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/427/original/590979.png" group-title="Спорт",Eurosport 2 HD
http://playlist.tv.planeta.tc/playlist/hls/427-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="834" tvg-name="Russian Extreme HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/834/original/6434388.png" group-title="Спорт",Russian Extreme HD
http://playlist.tv.planeta.tc/playlist/hls/834-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1234" tvg-name="Russian Extreme TV Ultra HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1234/original/7463148.png" group-title="Спорт",Russian Extreme TV Ultra HD
http://playlist.tv.planeta.tc/playlist/hls/1234-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="146" tvg-name="Extreme Sports Channel" tvg-logo="http://gl.weburg.net/00/tv/channels/1/146/original/594785.png" group-title="Спорт",Extreme Sports Channel
http://playlist.tv.planeta.tc/playlist/hls/146-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1228" tvg-name="M-1 Глобал" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1228/original/7461715.png" group-title="Спорт",M-1 Глобал
http://playlist.tv.planeta.tc/playlist/hls/1228-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="42" tvg-name="Viasat Sport" tvg-logo="http://gl.weburg.net/00/tv/channels/1/42/original/595901.png" group-title="Спорт",Viasat Sport
http://playlist.tv.planeta.tc/playlist/hls/42-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="519" tvg-name="БОКС ТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/519/original/1249145.png" group-title="Спорт",БОКС ТВ
http://playlist.tv.planeta.tc/playlist/hls/519-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="177" tvg-name="КХЛ ТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/177/original/594775.png" group-title="Спорт",КХЛ ТВ
http://playlist.tv.planeta.tc/playlist/hls/177-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="124" tvg-name="Матч! Страна" tvg-logo="http://gl.weburg.net/00/tv/channels/1/124/original/596234.png" group-title="Спорт",Матч! Страна
http://playlist.tv.planeta.tc/playlist/hls/124-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="242" tvg-name="Deutsche Welle HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/242/original/591481.png" group-title="Новости",Deutsche Welle HD
http://playlist.tv.planeta.tc/playlist/hls/242-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="245" tvg-name="FRANCE 24 HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/245/original/591477.png" group-title="Новости",FRANCE 24 HD
http://playlist.tv.planeta.tc/playlist/hls/245-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="439" tvg-name="RT HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/439/original/603464.png" group-title="Новости",RT HD
http://playlist.tv.planeta.tc/playlist/hls/439-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="529" tvg-name="BBC World News" tvg-logo="http://gl.weburg.net/00/tv/channels/1/529/original/1925895.png" group-title="Новости",BBC World News
http://playlist.tv.planeta.tc/playlist/hls/529-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="260" tvg-name="EuroNews" tvg-logo="http://gl.weburg.net/00/tv/channels/1/260/original/591479.png" group-title="Новости",EuroNews
http://playlist.tv.planeta.tc/playlist/hls/260-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="134" tvg-name="Russia Today" tvg-logo="http://gl.weburg.net/00/tv/channels/1/134/original/591475.png" group-title="Новости",Russia Today
http://playlist.tv.planeta.tc/playlist/hls/134-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="435" tvg-name="Мир 24" tvg-logo="http://gl.weburg.net/00/tv/channels/1/435/original/599045.png" group-title="Новости",Мир 24
http://playlist.tv.planeta.tc/playlist/hls/435-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="15" tvg-name="РБК ТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/15/original/591487.png" group-title="Новости",РБК ТВ
http://playlist.tv.planeta.tc/playlist/hls/15-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="491" tvg-name="360 HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/491/original/7459394.png" group-title="Региональные",360 HD
http://playlist.tv.planeta.tc/playlist/hls/491-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1275" tvg-name="РИМ HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1275/original/7497090.png" group-title="Региональные",РИМ HD
http://playlist.tv.planeta.tc/playlist/hls/1275-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1298" tvg-name="Каменск-ТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1298/original/7498187.png" group-title="Региональные",Каменск-ТВ
http://playlist.tv.planeta.tc/playlist/hls/1298-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1187" tvg-name="Контрольная закупка" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1187/original/7421193.png" group-title="Региональные",Контрольная закупка
http://playlist.tv.planeta.tc/playlist/hls/1187-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="481" tvg-name="КРИК - ТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/481/original/667629.png" group-title="Региональные",КРИК - ТВ
http://playlist.tv.planeta.tc/playlist/hls/481-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="13" tvg-name="ОТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/13/original/593823.png" group-title="Региональные",ОТВ
http://playlist.tv.planeta.tc/playlist/hls/13-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1340" tvg-name="Реальный Тагил" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1340/original/7507036.png" group-title="Региональные",Реальный Тагил
http://playlist.tv.planeta.tc/playlist/hls/1340-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="845" tvg-name="Смайл-ТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/845/original/6857408.png" group-title="Региональные",Смайл-ТВ
http://playlist.tv.planeta.tc/playlist/hls/845-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="63" tvg-name="Союз" tvg-logo="http://gl.weburg.net/00/tv/channels/1/63/original/591118.png" group-title="Региональные",Союз
http://playlist.tv.planeta.tc/playlist/hls/63-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="70" tvg-name="Студия 41" tvg-logo="http://gl.weburg.net/00/tv/channels/1/70/original/595229.png" group-title="Региональные",Студия 41
http://playlist.tv.planeta.tc/playlist/hls/70-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="283" tvg-name="Тагил-ТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/283/original/5733390.png" group-title="Региональные",Тагил-ТВ
http://playlist.tv.planeta.tc/playlist/hls/283-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="253" tvg-name="ТНВ Татарстан" tvg-logo="http://gl.weburg.net/00/tv/channels/1/253/original/595221.png" group-title="Региональные",ТНВ Татарстан
http://playlist.tv.planeta.tc/playlist/hls/253-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="21" tvg-name="Четвертый канал" tvg-logo="http://gl.weburg.net/00/tv/channels/1/21/original/591128.png" group-title="Региональные",Четвертый канал
http://playlist.tv.planeta.tc/playlist/hls/21-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="633" tvg-name="Эра-ТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/633/original/5662028.png" group-title="Региональные",Эра-ТВ
http://playlist.tv.planeta.tc/playlist/hls/633-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1276" tvg-name="Ювелирoчка" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1276/original/7497373.png" group-title="Региональные",Ювелирoчка
http://playlist.tv.planeta.tc/playlist/hls/1276-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="628" tvg-name="Матч ТВ HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/628/original/5519647.png" group-title="Федеральные",Матч ТВ HD
http://playlist.tv.planeta.tc/playlist/hls/628-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="420" tvg-name="Первый канал HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/420/original/591310.png" group-title="Федеральные",Первый канал HD
http://playlist.tv.planeta.tc/playlist/hls/420-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="419" tvg-name="Россия HD" tvg-logo="http://gl.weburg.net/00/tv/channels/1/419/original/590499.png" group-title="Федеральные",Россия HD
http://playlist.tv.planeta.tc/playlist/hls/419-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="1256" tvg-name="Домашний" tvg-logo="http://gl.weburg.net/00/tv/channels/1/1256/original/7486723.png" group-title="Федеральные",Домашний
http://playlist.tv.planeta.tc/playlist/hls/1256-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="259" tvg-name="Звезда" tvg-logo="http://gl.weburg.net/00/tv/channels/1/259/original/591096.png" group-title="Федеральные",Звезда
http://playlist.tv.planeta.tc/playlist/hls/259-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="319" tvg-name="Матч ТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/319/original/591112.png" group-title="Федеральные",Матч ТВ
http://playlist.tv.planeta.tc/playlist/hls/319-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="88" tvg-name="Мир" tvg-logo="http://gl.weburg.net/00/tv/channels/1/88/original/591469.png" group-title="Федеральные",Мир
http://playlist.tv.planeta.tc/playlist/hls/88-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="441" tvg-name="МУЗ - ТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/441/original/604344.png" group-title="Федеральные",МУЗ - ТВ
http://playlist.tv.planeta.tc/playlist/hls/441-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="12" tvg-name="НТВ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/12/original/591098.png" group-title="Федеральные",НТВ
http://playlist.tv.planeta.tc/playlist/hls/12-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="437" tvg-name="ОТР" tvg-logo="http://gl.weburg.net/00/tv/channels/1/437/original/603469.png" group-title="Федеральные",ОТР
http://playlist.tv.planeta.tc/playlist/hls/437-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="14" tvg-name="Первый канал" tvg-logo="http://gl.weburg.net/00/tv/channels/1/14/original/590496.png" group-title="Федеральные",Первый канал
http://playlist.tv.planeta.tc/playlist/hls/14-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="3" tvg-name="Пятый канал" tvg-logo="http://gl.weburg.net/00/tv/channels/1/3/original/591104.png" group-title="Федеральные",Пятый канал
http://playlist.tv.planeta.tc/playlist/hls/3-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="6" tvg-name="РЕН-Урал" tvg-logo="http://gl.weburg.net/00/tv/channels/1/6/original/591106.png" group-title="Федеральные",РЕН-Урал
http://playlist.tv.planeta.tc/playlist/hls/6-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="16" tvg-name="Россия 1" tvg-logo="http://gl.weburg.net/00/tv/channels/1/16/original/591108.png" group-title="Федеральные",Россия 1
http://playlist.tv.planeta.tc/playlist/hls/16-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="240" tvg-name="Россия 24" tvg-logo="http://gl.weburg.net/00/tv/channels/1/240/original/591114.png" group-title="Федеральные",Россия 24
http://playlist.tv.planeta.tc/playlist/hls/240-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="11" tvg-name="Россия К" tvg-logo="http://gl.weburg.net/00/tv/channels/1/11/original/591116.png" group-title="Федеральные",Россия К
http://playlist.tv.planeta.tc/playlist/hls/11-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="614" tvg-name="Спас" tvg-logo="http://gl.weburg.net/00/tv/channels/1/614/original/4885513.png" group-title="Федеральные",Спас
http://playlist.tv.planeta.tc/playlist/hls/614-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="17" tvg-name="СТС-Урал" tvg-logo="http://gl.weburg.net/00/tv/channels/1/17/original/591120.png" group-title="Федеральные",СТС-Урал
http://playlist.tv.planeta.tc/playlist/hls/17-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="18" tvg-name="ТВ Центр" tvg-logo="http://gl.weburg.net/00/tv/channels/1/18/original/591122.png" group-title="Федеральные",ТВ Центр
http://playlist.tv.planeta.tc/playlist/hls/18-live-0-master.m3u8?quality=min&videoOnly=true
#EXTINF:-1 tvg-id="19" tvg-name="ТНТ" tvg-logo="http://gl.weburg.net/00/tv/channels/1/19/original/591124.png" group-title="Федеральные",ТНТ
http://playlist.tv.planeta.tc/playlist/hls/19-live-0-master.m3u8?quality=min&videoOnly=true
`)

var wr bytes.Buffer

func TestM3uFile(t *testing.T) {
	m3u(httptest.NewRecorder(), "localhost")
}

func BenchmarkM3uFile(B *testing.B) {
	for i := 0; i < B.N; i++ {
		m3u(httptest.NewRecorder(), "localhost")
	}
}

func TestBasicAuth(t *testing.T) {
	changeBody(&wr, rd, []byte("localhost/"))
}

func BenchmarkChangeBody(B *testing.B) {
	for i := 0; i < B.N; i++ {
		changeBody(&wr, rd, []byte("localhost/"))
	}
}
