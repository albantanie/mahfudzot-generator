package database

import (
	"fmt"
	"log"

	"github.com/albantanie/mahfudzot-generator/internal/models"
)

// GetSeedData returns a comprehensive collection of mahfudzot (Arabic quotes)
func GetSeedData() []*models.QuoteRequest {
	return []*models.QuoteRequest{
		// Quotes from Prophet Muhammad (SAW)
		{
			TextArabic:  "إنما الأعمال بالنيات",
			TextLatin:   "Innama al-a'malu bin-niyyat",
			Translation: "Actions are but by intention",
			Author:      "Prophet Muhammad",
			Category:    "Intention",
			Source:      "Sahih Bukhari",
		},
		{
			TextArabic:  "اطلبوا العلم من المهد إلى اللحد",
			TextLatin:   "Utlubu al-'ilma min al-mahdi ila al-lahd",
			Translation: "Seek knowledge from the cradle to the grave",
			Author:      "Prophet Muhammad",
			Category:    "Knowledge",
			Source:      "Hadith",
		},
		{
			TextArabic:  "الصبر مفتاح الفرج",
			TextLatin:   "As-sabru miftahu al-faraj",
			Translation: "Patience is the key to relief",
			Author:      "Prophet Muhammad",
			Category:    "Patience",
			Source:      "Hadith",
		},
		{
			TextArabic:  "من كان في حاجة أخيه كان الله في حاجته",
			TextLatin:   "Man kana fi hajati akhihi kana Allahu fi hajatih",
			Translation: "Whoever helps his brother, Allah will help him",
			Author:      "Prophet Muhammad",
			Category:    "Brotherhood",
			Source:      "Sahih Bukhari",
		},
		{
			TextArabic:  "خير الناس أنفعهم للناس",
			TextLatin:   "Khairu an-nasi anfa'uhum lin-nas",
			Translation: "The best of people are those who benefit others",
			Author:      "Prophet Muhammad",
			Category:    "Service",
			Source:      "Hadith",
		},

		// Quotes from Imam Ali (RA)
		{
			TextArabic:  "العلم نور",
			TextLatin:   "Al-'ilmu nur",
			Translation: "Knowledge is light",
			Author:      "Imam Ali",
			Category:    "Knowledge",
			Source:      "Nahj al-Balagha",
		},
		{
			TextArabic:  "الدنيا دار ممر لا دار مقر",
			TextLatin:   "Ad-dunya daru mamarrin la daru muqarr",
			Translation: "This world is a place of passage, not a place of residence",
			Author:      "Imam Ali",
			Category:    "Wisdom",
			Source:      "Nahj al-Balagha",
		},
		{
			TextArabic:  "من عرف نفسه فقد عرف ربه",
			TextLatin:   "Man 'arafa nafsahu faqad 'arafa rabbah",
			Translation: "Whoever knows himself knows his Lord",
			Author:      "Imam Ali",
			Category:    "Self-Knowledge",
			Source:      "Nahj al-Balagha",
		},
		{
			TextArabic:  "الصمت حكمة وقليل فاعله",
			TextLatin:   "As-samtu hikmah wa qalilun fa'iluh",
			Translation: "Silence is wisdom, but few practice it",
			Author:      "Imam Ali",
			Category:    "Wisdom",
			Source:      "Nahj al-Balagha",
		},
		{
			TextArabic:  "لا تكن عبداً لغيرك وقد جعلك الله حراً",
			TextLatin:   "La takun 'abdan li-ghayriki wa qad ja'alaka Allahu hurran",
			Translation: "Do not be a slave to others when Allah has made you free",
			Author:      "Imam Ali",
			Category:    "Freedom",
			Source:      "Nahj al-Balagha",
		},

		// Quotes from Imam Al-Ghazali
		{
			TextArabic:  "العلم ما نفع ليس العلم ما حفظ",
			TextLatin:   "Al-'ilmu ma nafa'a laysa al-'ilmu ma hufiza",
			Translation: "Knowledge is what benefits, not what is memorized",
			Author:      "Imam Al-Ghazali",
			Category:    "Knowledge",
			Source:      "Ihya Ulum al-Din",
		},
		{
			TextArabic:  "القلب إذا أقبل على الله أقبل الله عليه",
			TextLatin:   "Al-qalbu idha aqbala 'ala Allah aqbala Allahu 'alayh",
			Translation: "When the heart turns to Allah, Allah turns to it",
			Author:      "Imam Al-Ghazali",
			Category:    "Spirituality",
			Source:      "Ihya Ulum al-Din",
		},
		{
			TextArabic:  "الدنيا مزرعة الآخرة",
			TextLatin:   "Ad-dunya mazra'atu al-akhirah",
			Translation: "This world is the farm of the hereafter",
			Author:      "Imam Al-Ghazali",
			Category:    "Life",
			Source:      "Ihya Ulum al-Din",
		},

		// Quotes from Ibn Sina (Avicenna)
		{
			TextArabic:  "الجهل موت الأحياء",
			TextLatin:   "Al-jahlu mawtu al-ahya'",
			Translation: "Ignorance is the death of the living",
			Author:      "Ibn Sina",
			Category:    "Knowledge",
			Source:      "Al-Qanun fi al-Tibb",
		},
		{
			TextArabic:  "العقل السليم في الجسم السليم",
			TextLatin:   "Al-'aqlu as-salimu fi al-jismi as-salim",
			Translation: "A sound mind in a sound body",
			Author:      "Ibn Sina",
			Category:    "Health",
			Source:      "Medical Works",
		},

		// Quotes from Al-Mutanabbi
		{
			TextArabic:  "من طلب العلا سهر الليالي",
			TextLatin:   "Man talaba al-'ula sahira al-layali",
			Translation: "Whoever seeks excellence stays awake at night",
			Author:      "Al-Mutanabbi",
			Category:    "Excellence",
			Source:      "Diwan Al-Mutanabbi",
		},
		{
			TextArabic:  "على قدر أهل العزم تأتي العزائم",
			TextLatin:   "Ala qadri ahli al-'azmi ta'ti al-'aza'im",
			Translation: "Great deeds come from people of great determination",
			Author:      "Al-Mutanabbi",
			Category:    "Determination",
			Source:      "Diwan Al-Mutanabbi",
		},
		{
			TextArabic:  "ومن يك ذا فم مر مريض يجد مراً به الماء الزلالا",
			TextLatin:   "Wa man yaku dha famin murrin maridin yajid murran bihi al-ma'a az-zulala",
			Translation: "One with a bitter sick mouth will find even pure water bitter",
			Author:      "Al-Mutanabbi",
			Category:    "Perspective",
			Source:      "Diwan Al-Mutanabbi",
		},

		// Quotes from Ibn Khaldun
		{
			TextArabic:  "العصبية أساس الملك",
			TextLatin:   "Al-'asabiyyatu asasu al-mulk",
			Translation: "Social cohesion is the foundation of power",
			Author:      "Ibn Khaldun",
			Category:    "Society",
			Source:      "Al-Muqaddimah",
		},
		{
			TextArabic:  "التاريخ في ظاهره لا يزيد عن الإخبار",
			TextLatin:   "At-tarikhu fi zahirihi la yazidu 'an al-ikhbar",
			Translation: "History on its surface is nothing more than information",
			Author:      "Ibn Khaldun",
			Category:    "History",
			Source:      "Al-Muqaddimah",
		},

		// Quotes from Imam Ash-Shafi'i
		{
			TextArabic:  "ما جادلت أحداً إلا تمنيت أن يظهر الله الحق على لسانه",
			TextLatin:   "Ma jadaltu ahadan illa tamannaytu an yuzhira Allahu al-haqqa 'ala lisanih",
			Translation: "I never debated anyone except I wished Allah would show the truth through their tongue",
			Author:      "Imam Ash-Shafi'i",
			Category:    "Humility",
			Source:      "Manaqib Ash-Shafi'i",
		},
		{
			TextArabic:  "كلما ازددت علماً ازددت علماً بجهلي",
			TextLatin:   "Kullama izdadtu 'ilman izdadtu 'ilman bi-jahli",
			Translation: "The more I learn, the more I realize my ignorance",
			Author:      "Imam Ash-Shafi'i",
			Category:    "Humility",
			Source:      "Sayings",
		},

		// Quotes from Ibn Taymiyyah
		{
			TextArabic:  "القلب لا يستقيم إلا بالتوحيد",
			TextLatin:   "Al-qalbu la yastaqimu illa bit-tawhid",
			Translation: "The heart cannot be upright except through monotheism",
			Author:      "Ibn Taymiyyah",
			Category:    "Faith",
			Source:      "Majmu' al-Fatawa",
		},
		{
			TextArabic:  "من أراد السعادة الأبدية فليلزم عتبة العبودية",
			TextLatin:   "Man arada as-sa'adat al-abadiyyata falyalzam 'atabat al-'ubudiyyah",
			Translation: "Whoever wants eternal happiness should stick to the threshold of servitude",
			Author:      "Ibn Taymiyyah",
			Category:    "Spirituality",
			Source:      "Al-Ubudiyyah",
		},

		// Quotes from Al-Jahiz
		{
			TextArabic:  "الكتاب أستاذ لا يعنف ومعلم لا يغضب",
			TextLatin:   "Al-kitabu ustadhun la yu'annifu wa mu'allimun la yaghdhab",
			Translation: "A book is a teacher that doesn't scold and an instructor that doesn't get angry",
			Author:      "Al-Jahiz",
			Category:    "Knowledge",
			Source:      "Al-Bayan wa al-Tabyin",
		},

		// Quotes from Ibn Rushd (Averroes)
		{
			TextArabic:  "الجهل يؤدي إلى الخوف والخوف يؤدي إلى الكراهية",
			TextLatin:   "Al-jahlu yu'addi ila al-khawfi wal-khawfu yu'addi ila al-karahiyyah",
			Translation: "Ignorance leads to fear, and fear leads to hatred",
			Author:      "Ibn Rushd",
			Category:    "Wisdom",
			Source:      "Philosophical Works",
		},

		// Quotes from Al-Kindi
		{
			TextArabic:  "لا نستحي من قول الحق واقتباس الحق من أين أتى",
			TextLatin:   "La nastahi min qawli al-haqqi waqtibasi al-haqqi min ayna ata",
			Translation: "We should not be ashamed to speak the truth and acquire truth from wherever it comes",
			Author:      "Al-Kindi",
			Category:    "Truth",
			Source:      "Philosophical Treatises",
		},

		// Quotes from Imam Ahmad ibn Hanbal
		{
			TextArabic:  "العلم لا يعطيك بعضه حتى تعطيه كلك",
			TextLatin:   "Al-'ilmu la yu'tika ba'dhahu hatta tu'tiyahu kullak",
			Translation: "Knowledge will not give you part of it until you give it all of yourself",
			Author:      "Imam Ahmad ibn Hanbal",
			Category:    "Knowledge",
			Source:      "Sayings",
		},
		{
			TextArabic:  "الناس إلى العدل أحوج منهم إلى الماء والنار",
			TextLatin:   "An-nasu ila al-'adli ahwaju minhum ila al-ma'i wan-nar",
			Translation: "People need justice more than they need water and fire",
			Author:      "Imam Ahmad ibn Hanbal",
			Category:    "Justice",
			Source:      "Musnad Ahmad",
		},

		// Quotes from Al-Farabi
		{
			TextArabic:  "الفضيلة وسط بين رذيلتين",
			TextLatin:   "Al-fadilatu wasatun bayna radhilatayn",
			Translation: "Virtue is the middle path between two vices",
			Author:      "Al-Farabi",
			Category:    "Ethics",
			Source:      "Al-Madina al-Fadila",
		},
		{
			TextArabic:  "السعادة هي الخير الأعظم",
			TextLatin:   "As-sa'adatu hiya al-khayru al-a'zam",
			Translation: "Happiness is the greatest good",
			Author:      "Al-Farabi",
			Category:    "Happiness",
			Source:      "Tahsil al-Sa'ada",
		},

		// Quotes from Ibn al-Qayyim
		{
			TextArabic:  "القلوب آنية الله في أرضه",
			TextLatin:   "Al-qulubu aniyatu Allahi fi ardhih",
			Translation: "Hearts are Allah's vessels on His earth",
			Author:      "Ibn al-Qayyim",
			Category:    "Spirituality",
			Source:      "Madarij al-Salikin",
		},
		{
			TextArabic:  "الدعاء مخ العبادة",
			TextLatin:   "Ad-du'a'u mukhkhu al-'ibadah",
			Translation: "Prayer is the essence of worship",
			Author:      "Ibn al-Qayyim",
			Category:    "Prayer",
			Source:      "Al-Jawab al-Kafi",
		},

		// Quotes from Al-Razi (Fakhr al-Din)
		{
			TextArabic:  "العقل نور والنقل نور ولا تعارض بين نورين",
			TextLatin:   "Al-'aqlu nurun wan-naqlu nurun wa la ta'aruda bayna nurayn",
			Translation: "Reason is light and revelation is light, and there is no contradiction between two lights",
			Author:      "Fakhr al-Din al-Razi",
			Category:    "Reason",
			Source:      "Mafatih al-Ghayb",
		},

		// Quotes from Ibn Arabi
		{
			TextArabic:  "من عرف نفسه عرف ربه",
			TextLatin:   "Man 'arafa nafsahu 'arafa rabbah",
			Translation: "Whoever knows himself knows his Lord",
			Author:      "Ibn Arabi",
			Category:    "Self-Knowledge",
			Source:      "Fusus al-Hikam",
		},
		{
			TextArabic:  "الكون كله كتاب الله المنشور",
			TextLatin:   "Al-kawnu kulluhu kitabu Allahi al-manshur",
			Translation: "The entire universe is Allah's open book",
			Author:      "Ibn Arabi",
			Category:    "Universe",
			Source:      "Al-Futuhat al-Makkiyyah",
		},

		// Quotes from Al-Bukhari
		{
			TextArabic:  "ما كتبت حديثاً إلا اغتسلت قبله وصليت ركعتين",
			TextLatin:   "Ma katabtu hadithan illa ightasaltu qablahu wa sallaytu rak'atayn",
			Translation: "I never wrote a hadith except that I performed ablution before it and prayed two units",
			Author:      "Imam Al-Bukhari",
			Category:    "Scholarship",
			Source:      "Biography",
		},

		// Quotes from Muslim ibn al-Hajjaj
		{
			TextArabic:  "الإسناد من الدين ولولا الإسناد لقال من شاء ما شاء",
			TextLatin:   "Al-isnadu min ad-dini wa lawla al-isnadu laqala man sha'a ma sha'a",
			Translation: "Chain of narration is part of religion; without it, anyone could say whatever they wanted",
			Author:      "Imam Muslim",
			Category:    "Scholarship",
			Source:      "Sahih Muslim Introduction",
		},

		// Quotes from Al-Nawawi
		{
			TextArabic:  "من سلك طريقاً يلتمس فيه علماً سهل الله له طريقاً إلى الجنة",
			TextLatin:   "Man salaka tariqan yaltamisu fihi 'ilman sahhal Allahu lahu tariqan ila al-jannah",
			Translation: "Whoever travels a path seeking knowledge, Allah will make easy for him a path to Paradise",
			Author:      "Imam An-Nawawi",
			Category:    "Knowledge",
			Source:      "Riyadh as-Salihin",
		},

		// Quotes from Ibn Kathir
		{
			TextArabic:  "القرآن يفسر بعضه بعضاً",
			TextLatin:   "Al-Qur'anu yufassiru ba'duhu ba'dan",
			Translation: "The Quran explains parts of itself through other parts",
			Author:      "Ibn Kathir",
			Category:    "Quran",
			Source:      "Tafsir Ibn Kathir",
		},

		// Quotes from Al-Tabari
		{
			TextArabic:  "لا يستغني طالب العلم عن أربعة: ذكاء الطبع وطول الباع وكثرة الاطلاع وطول العمر",
			TextLatin:   "La yastaghni talibu al-'ilmi 'an arba'ah: dhaka'u at-tab'i wa tulu al-ba'i wa kathratu al-ittila'i wa tulu al-'umr",
			Translation: "A seeker of knowledge cannot do without four things: natural intelligence, extensive reach, broad reading, and long life",
			Author:      "Al-Tabari",
			Category:    "Knowledge",
			Source:      "Tafsir al-Tabari",
		},

		// Quotes from Al-Qurtubi
		{
			TextArabic:  "العبرة بعموم اللفظ لا بخصوص السبب",
			TextLatin:   "Al-'ibratu bi-'umumi al-lafzi la bi-khususi as-sabab",
			Translation: "Consideration is given to the generality of the wording, not the specificity of the reason",
			Author:      "Al-Qurtubi",
			Category:    "Jurisprudence",
			Source:      "Tafsir al-Qurtubi",
		},

		// Quotes from Ibn Hazm
		{
			TextArabic:  "من أراد أن ينصف من نفسه فليتوهم نفسه خصماً ومن خالفه منصفاً",
			TextLatin:   "Man arada an yunsifa min nafsihi falyatawahham nafsahu khasman wa man khalafahu munsifan",
			Translation: "Whoever wants to be fair to himself should imagine himself as an opponent and his opponent as fair",
			Author:      "Ibn Hazm",
			Category:    "Justice",
			Source:      "Al-Akhlaq wa al-Siyar",
		},
		{
			TextArabic:  "آفة العلماء الوقوف مع المتشابه",
			TextLatin:   "Afatu al-'ulama'i al-wuqufu ma'a al-mutashabih",
			Translation: "The bane of scholars is stopping at ambiguous matters",
			Author:      "Ibn Hazm",
			Category:    "Scholarship",
			Source:      "Al-Ihkam fi Usul al-Ahkam",
		},

		// Quotes from Imam Malik
		{
			TextArabic:  "ما منا إلا راد ومردود عليه إلا صاحب هذا القبر",
			TextLatin:   "Ma minna illa raddun wa mardudun 'alayhi illa sahibu hadha al-qabr",
			Translation: "None of us is free from error and being corrected, except the occupant of this grave (Prophet Muhammad)",
			Author:      "Imam Malik",
			Category:    "Humility",
			Source:      "Al-Muwatta",
		},
		{
			TextArabic:  "لن يصلح آخر هذه الأمة إلا بما صلح به أولها",
			TextLatin:   "Lan yasluh akhiru hadhihi al-ummati illa bima salaha bihi awwaluha",
			Translation: "The latter part of this nation will not be reformed except by that which reformed its early part",
			Author:      "Imam Malik",
			Category:    "Reform",
			Source:      "Sayings",
		},

		// Quotes from Abu Hanifa
		{
			TextArabic:  "لولا السنتان لهلك النعمان",
			TextLatin:   "Lawla as-sanatan lahalaka an-Nu'man",
			Translation: "Were it not for the two years (with Abu Hanifa's teachers), Nu'man would have perished",
			Author:      "Imam Abu Hanifa",
			Category:    "Learning",
			Source:      "Biography",
		},
		{
			TextArabic:  "الفقه أفضل من العبادة",
			TextLatin:   "Al-fiqhu afdalu min al-'ibadah",
			Translation: "Understanding (jurisprudence) is better than worship",
			Author:      "Imam Abu Hanifa",
			Category:    "Knowledge",
			Source:      "Sayings",
		},

		// Quotes from Al-Junayd
		{
			TextArabic:  "التصوف أن تكون مع الله بلا علاقة",
			TextLatin:   "At-tasawwufu an takuna ma'a Allahi bila 'alaqah",
			Translation: "Sufism is to be with Allah without attachment",
			Author:      "Al-Junayd",
			Category:    "Spirituality",
			Source:      "Sufi Teachings",
		},
		{
			TextArabic:  "الطرق إلى الله بعدد أنفاس الخلائق",
			TextLatin:   "At-turuqu ila Allahi bi-'adadi anfasi al-khala'iq",
			Translation: "The paths to Allah are as numerous as the breaths of creation",
			Author:      "Al-Junayd",
			Category:    "Spirituality",
			Source:      "Sufi Teachings",
		},

		// Quotes from Al-Hallaj
		{
			TextArabic:  "من لم تحرقه المحبة فهو ناقص الوضوء",
			TextLatin:   "Man lam tuhriqhu al-mahabbatu fahuwa naqisu al-wudu'",
			Translation: "Whoever is not burned by love has incomplete ablution",
			Author:      "Al-Hallaj",
			Category:    "Love",
			Source:      "Diwan al-Hallaj",
		},

		// Quotes from Rumi (Jalal ad-Din)
		{
			TextArabic:  "كن كالماء في التواضع وكالنار في الهمة",
			TextLatin:   "Kun kal-ma'i fi at-tawadu'i wa kan-nari fi al-himmah",
			Translation: "Be like water in humility and like fire in determination",
			Author:      "Rumi",
			Category:    "Character",
			Source:      "Masnavi",
		},
		{
			TextArabic:  "أمس ذهب وغداً لم يأت واليوم بين يديك",
			TextLatin:   "Amsi dhahaba wa ghadan lam ya'ti wal-yawmu bayna yadayk",
			Translation: "Yesterday is gone, tomorrow has not come, and today is in your hands",
			Author:      "Rumi",
			Category:    "Time",
			Source:      "Masnavi",
		},

		// Quotes from Saadi Shirazi
		{
			TextArabic:  "بني آدم أعضاء جسد واحد",
			TextLatin:   "Bani Adama a'da'u jasadin wahid",
			Translation: "Human beings are members of one body",
			Author:      "Saadi Shirazi",
			Category:    "Humanity",
			Source:      "Gulistan",
		},

		// Quotes from Hafez
		{
			TextArabic:  "لا تحزن إن لم تفهم أسرار الحب",
			TextLatin:   "La tahzan in lam tafham asrara al-hubb",
			Translation: "Do not grieve if you do not understand the secrets of love",
			Author:      "Hafez",
			Category:    "Love",
			Source:      "Diwan Hafez",
		},

		// Quotes from Omar Khayyam
		{
			TextArabic:  "اشرب الخمر واترك الحكمة للحكماء",
			TextLatin:   "Ishrab al-khamra watruk al-hikmata lil-hukama'",
			Translation: "Drink wine and leave wisdom to the wise",
			Author:      "Omar Khayyam",
			Category:    "Philosophy",
			Source:      "Rubaiyat",
		},

		// Quotes from Al-Biruni
		{
			TextArabic:  "العلم أشرف ما رغب فيه الراغب",
			TextLatin:   "Al-'ilmu ashrafu ma raghiba fihi ar-raghib",
			Translation: "Knowledge is the noblest thing a seeker can desire",
			Author:      "Al-Biruni",
			Category:    "Knowledge",
			Source:      "Scientific Works",
		},

		// Quotes from Ibn Battuta
		{
			TextArabic:  "السفر يعلم الصبر",
			TextLatin:   "As-safaru yu'allimu as-sabr",
			Translation: "Travel teaches patience",
			Author:      "Ibn Battuta",
			Category:    "Travel",
			Source:      "Rihla",
		},

		// Quotes from Al-Mas'udi
		{
			TextArabic:  "التاريخ مرآة الأمم",
			TextLatin:   "At-tarikhu mir'atu al-umam",
			Translation: "History is the mirror of nations",
			Author:      "Al-Mas'udi",
			Category:    "History",
			Source:      "Muruj adh-Dhahab",
		},

		// Quotes from Ibn al-Athir
		{
			TextArabic:  "العدل أساس الملك",
			TextLatin:   "Al-'adlu asasu al-mulk",
			Translation: "Justice is the foundation of rule",
			Author:      "Ibn al-Athir",
			Category:    "Justice",
			Source:      "Al-Kamil fi at-Tarikh",
		},

		// Quotes from Al-Suyuti
		{
			TextArabic:  "طلب العلم فريضة على كل مسلم ومسلمة",
			TextLatin:   "Talabu al-'ilmi faridatun 'ala kulli muslimin wa muslimah",
			Translation: "Seeking knowledge is an obligation upon every Muslim man and woman",
			Author:      "Al-Suyuti",
			Category:    "Knowledge",
			Source:      "Jami' as-Saghir",
		},

		// Quotes from Ibn Qudamah
		{
			TextArabic:  "من استوى يوماه فهو مغبون",
			TextLatin:   "Man istawaya yawmahu fahuwa maghbun",
			Translation: "Whoever's two days are equal is at a loss",
			Author:      "Ibn Qudamah",
			Category:    "Progress",
			Source:      "Minhaj al-Qasidin",
		},

		// Quotes from Al-Dhahabi
		{
			TextArabic:  "العلم نور والعمل نور ونور على نور",
			TextLatin:   "Al-'ilmu nurun wal-'amalu nurun wa nurun 'ala nur",
			Translation: "Knowledge is light, action is light, and light upon light",
			Author:      "Al-Dhahabi",
			Category:    "Knowledge",
			Source:      "Siyar A'lam an-Nubala",
		},

		// Additional wisdom quotes
		{
			TextArabic:  "من صبر ظفر",
			TextLatin:   "Man sabara zafar",
			Translation: "Whoever is patient will triumph",
			Author:      "Arabic Proverb",
			Category:    "Patience",
			Source:      "Traditional Wisdom",
		},
		{
			TextArabic:  "العقل زينة والجهل شين",
			TextLatin:   "Al-'aqlu zinatun wal-jahlu shayn",
			Translation: "Intelligence is an ornament and ignorance is a disgrace",
			Author:      "Arabic Proverb",
			Category:    "Wisdom",
			Source:      "Traditional Wisdom",
		},
		{
			TextArabic:  "من جد وجد ومن زرع حصد",
			TextLatin:   "Man jadda wajada wa man zara'a hasad",
			Translation: "Whoever strives will find, and whoever sows will reap",
			Author:      "Arabic Proverb",
			Category:    "Effort",
			Source:      "Traditional Wisdom",
		},
		{
			TextArabic:  "الصديق وقت الضيق",
			TextLatin:   "As-sadiqu waqtu ad-diq",
			Translation: "A friend in need is a friend indeed",
			Author:      "Arabic Proverb",
			Category:    "Friendship",
			Source:      "Traditional Wisdom",
		},
		{
			TextArabic:  "درهم وقاية خير من قنطار علاج",
			TextLatin:   "Dirhamu wiqayatin khayrun min qintari 'ilaj",
			Translation: "An ounce of prevention is worth a pound of cure",
			Author:      "Arabic Proverb",
			Category:    "Prevention",
			Source:      "Traditional Wisdom",
		},
	}
}

// SeedDatabase populates the database with initial quote data
func SeedDatabase(db QuoteRepository) error {
	quotes := GetSeedData()

	log.Printf("Starting to seed database with %d quotes...", len(quotes))

	successCount := 0
	for i, quote := range quotes {
		_, err := db.Create(quote)
		if err != nil {
			log.Printf("Failed to insert quote %d: %v", i+1, err)
			continue
		}
		successCount++
	}

	log.Printf("Successfully seeded %d out of %d quotes", successCount, len(quotes))

	if successCount == 0 {
		return fmt.Errorf("failed to seed any quotes")
	}

	return nil
}

// GetQuoteCount returns the total number of quotes in seed data
func GetQuoteCount() int {
	return len(GetSeedData())
}
