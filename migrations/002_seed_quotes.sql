-- Seed data for mahfudzot quotes
-- This file contains comprehensive Arabic quotes from various Islamic scholars

INSERT INTO quotes (text_arabic, text_latin, translation, author, category, source) VALUES
-- Prophet Muhammad (SAW) quotes
('إنما الأعمال بالنيات', 'Innama al-a''malu bin-niyyat', 'Actions are but by intention', 'Prophet Muhammad', 'Intention', 'Sahih Bukhari'),
('اطلبوا العلم من المهد إلى اللحد', 'Utlubu al-''ilma min al-mahdi ila al-lahd', 'Seek knowledge from the cradle to the grave', 'Prophet Muhammad', 'Knowledge', 'Hadith'),
('الصبر مفتاح الفرج', 'As-sabru miftahu al-faraj', 'Patience is the key to relief', 'Prophet Muhammad', 'Patience', 'Hadith'),
('من كان في حاجة أخيه كان الله في حاجته', 'Man kana fi hajati akhihi kana Allahu fi hajatih', 'Whoever helps his brother, Allah will help him', 'Prophet Muhammad', 'Brotherhood', 'Sahih Bukhari'),
('خير الناس أنفعهم للناس', 'Khairu an-nasi anfa''uhum lin-nas', 'The best of people are those who benefit others', 'Prophet Muhammad', 'Service', 'Hadith'),

-- Imam Ali (RA) quotes
('العلم نور', 'Al-''ilmu nur', 'Knowledge is light', 'Imam Ali', 'Knowledge', 'Nahj al-Balagha'),
('الدنيا دار ممر لا دار مقر', 'Ad-dunya daru mamarrin la daru muqarr', 'This world is a place of passage, not a place of residence', 'Imam Ali', 'Wisdom', 'Nahj al-Balagha'),
('من عرف نفسه فقد عرف ربه', 'Man ''arafa nafsahu faqad ''arafa rabbah', 'Whoever knows himself knows his Lord', 'Imam Ali', 'Self-Knowledge', 'Nahj al-Balagha'),
('الصمت حكمة وقليل فاعله', 'As-samtu hikmah wa qalilun fa''iluh', 'Silence is wisdom, but few practice it', 'Imam Ali', 'Wisdom', 'Nahj al-Balagha'),
('لا تكن عبداً لغيرك وقد جعلك الله حراً', 'La takun ''abdan li-ghayriki wa qad ja''alaka Allahu hurran', 'Do not be a slave to others when Allah has made you free', 'Imam Ali', 'Freedom', 'Nahj al-Balagha'),

-- Imam Al-Ghazali quotes
('العلم ما نفع ليس العلم ما حفظ', 'Al-''ilmu ma nafa''a laysa al-''ilmu ma hufiza', 'Knowledge is what benefits, not what is memorized', 'Imam Al-Ghazali', 'Knowledge', 'Ihya Ulum al-Din'),
('القلب إذا أقبل على الله أقبل الله عليه', 'Al-qalbu idha aqbala ''ala Allah aqbala Allahu ''alayh', 'When the heart turns to Allah, Allah turns to it', 'Imam Al-Ghazali', 'Spirituality', 'Ihya Ulum al-Din'),
('الدنيا مزرعة الآخرة', 'Ad-dunya mazra''atu al-akhirah', 'This world is the farm of the hereafter', 'Imam Al-Ghazali', 'Life', 'Ihya Ulum al-Din'),

-- Ibn Sina quotes
('الجهل موت الأحياء', 'Al-jahlu mawtu al-ahya''', 'Ignorance is the death of the living', 'Ibn Sina', 'Knowledge', 'Al-Qanun fi al-Tibb'),
('العقل السليم في الجسم السليم', 'Al-''aqlu as-salimu fi al-jismi as-salim', 'A sound mind in a sound body', 'Ibn Sina', 'Health', 'Medical Works'),

-- Al-Mutanabbi quotes
('من طلب العلا سهر الليالي', 'Man talaba al-''ula sahira al-layali', 'Whoever seeks excellence stays awake at night', 'Al-Mutanabbi', 'Excellence', 'Diwan Al-Mutanabbi'),
('على قدر أهل العزم تأتي العزائم', 'Ala qadri ahli al-''azmi ta''ti al-''aza''im', 'Great deeds come from people of great determination', 'Al-Mutanabbi', 'Determination', 'Diwan Al-Mutanabbi'),
('ومن يك ذا فم مر مريض يجد مراً به الماء الزلالا', 'Wa man yaku dha famin murrin maridin yajid murran bihi al-ma''a az-zulala', 'One with a bitter sick mouth will find even pure water bitter', 'Al-Mutanabbi', 'Perspective', 'Diwan Al-Mutanabbi'),

-- Ibn Khaldun quotes
('العصبية أساس الملك', 'Al-''asabiyyatu asasu al-mulk', 'Social cohesion is the foundation of power', 'Ibn Khaldun', 'Society', 'Al-Muqaddimah'),
('التاريخ في ظاهره لا يزيد عن الإخبار', 'At-tarikhu fi zahirihi la yazidu ''an al-ikhbar', 'History on its surface is nothing more than information', 'Ibn Khaldun', 'History', 'Al-Muqaddimah'),

-- Imam Ash-Shafi'i quotes
('ما جادلت أحداً إلا تمنيت أن يظهر الله الحق على لسانه', 'Ma jadaltu ahadan illa tamannaytu an yuzhira Allahu al-haqqa ''ala lisanih', 'I never debated anyone except I wished Allah would show the truth through their tongue', 'Imam Ash-Shafi''i', 'Humility', 'Manaqib Ash-Shafi''i'),
('كلما ازددت علماً ازددت علماً بجهلي', 'Kullama izdadtu ''ilman izdadtu ''ilman bi-jahli', 'The more I learn, the more I realize my ignorance', 'Imam Ash-Shafi''i', 'Humility', 'Sayings'),

-- Ibn Taymiyyah quotes
('القلب لا يستقيم إلا بالتوحيد', 'Al-qalbu la yastaqimu illa bit-tawhid', 'The heart cannot be upright except through monotheism', 'Ibn Taymiyyah', 'Faith', 'Majmu'' al-Fatawa'),
('من أراد السعادة الأبدية فليلزم عتبة العبودية', 'Man arada as-sa''adat al-abadiyyata falyalzam ''atabat al-''ubudiyyah', 'Whoever wants eternal happiness should stick to the threshold of servitude', 'Ibn Taymiyyah', 'Spirituality', 'Al-Ubudiyyah'),

-- Al-Jahiz quotes
('الكتاب أستاذ لا يعنف ومعلم لا يغضب', 'Al-kitabu ustadhun la yu''annifu wa mu''allimun la yaghdhab', 'A book is a teacher that doesn''t scold and an instructor that doesn''t get angry', 'Al-Jahiz', 'Knowledge', 'Al-Bayan wa al-Tabyin'),

-- Ibn Rushd quotes
('الجهل يؤدي إلى الخوف والخوف يؤدي إلى الكراهية', 'Al-jahlu yu''addi ila al-khawfi wal-khawfu yu''addi ila al-karahiyyah', 'Ignorance leads to fear, and fear leads to hatred', 'Ibn Rushd', 'Wisdom', 'Philosophical Works'),

-- Al-Kindi quotes
('لا نستحي من قول الحق واقتباس الحق من أين أتى', 'La nastahi min qawli al-haqqi waqtibasi al-haqqi min ayna ata', 'We should not be ashamed to speak the truth and acquire truth from wherever it comes', 'Al-Kindi', 'Truth', 'Philosophical Treatises'),

-- Imam Ahmad ibn Hanbal quotes
('العلم لا يعطيك بعضه حتى تعطيه كلك', 'Al-''ilmu la yu''tika ba''dhahu hatta tu''tiyahu kullak', 'Knowledge will not give you part of it until you give it all of yourself', 'Imam Ahmad ibn Hanbal', 'Knowledge', 'Sayings'),
('الناس إلى العدل أحوج منهم إلى الماء والنار', 'An-nasu ila al-''adli ahwaju minhum ila al-ma''i wan-nar', 'People need justice more than they need water and fire', 'Imam Ahmad ibn Hanbal', 'Justice', 'Musnad Ahmad'),

-- Al-Farabi quotes
('الفضيلة وسط بين رذيلتين', 'Al-fadilatu wasatun bayna radhilatayn', 'Virtue is the middle path between two vices', 'Al-Farabi', 'Ethics', 'Al-Madina al-Fadila'),
('السعادة هي الخير الأعظم', 'As-sa''adatu hiya al-khayru al-a''zam', 'Happiness is the greatest good', 'Al-Farabi', 'Happiness', 'Tahsil al-Sa''ada'),

-- Ibn al-Qayyim quotes
('القلوب آنية الله في أرضه', 'Al-qulubu aniyatu Allahi fi ardhih', 'Hearts are Allah''s vessels on His earth', 'Ibn al-Qayyim', 'Spirituality', 'Madarij al-Salikin'),
('الدعاء مخ العبادة', 'Ad-du''a''u mukhkhu al-''ibadah', 'Prayer is the essence of worship', 'Ibn al-Qayyim', 'Prayer', 'Al-Jawab al-Kafi'),

-- Fakhr al-Din al-Razi quotes
('العقل نور والنقل نور ولا تعارض بين نورين', 'Al-''aqlu nurun wan-naqlu nurun wa la ta''aruda bayna nurayn', 'Reason is light and revelation is light, and there is no contradiction between two lights', 'Fakhr al-Din al-Razi', 'Reason', 'Mafatih al-Ghayb'),

-- Ibn Arabi quotes
('من عرف نفسه عرف ربه', 'Man ''arafa nafsahu ''arafa rabbah', 'Whoever knows himself knows his Lord', 'Ibn Arabi', 'Self-Knowledge', 'Fusus al-Hikam'),
('الكون كله كتاب الله المنشور', 'Al-kawnu kulluhu kitabu Allahi al-manshur', 'The entire universe is Allah''s open book', 'Ibn Arabi', 'Universe', 'Al-Futuhat al-Makkiyyah'),

-- Traditional wisdom
('من صبر ظفر', 'Man sabara zafar', 'Whoever is patient will triumph', 'Arabic Proverb', 'Patience', 'Traditional Wisdom'),
('العقل زينة والجهل شين', 'Al-''aqlu zinatun wal-jahlu shayn', 'Intelligence is an ornament and ignorance is a disgrace', 'Arabic Proverb', 'Wisdom', 'Traditional Wisdom'),
('من جد وجد ومن زرع حصد', 'Man jadda wajada wa man zara''a hasad', 'Whoever strives will find, and whoever sows will reap', 'Arabic Proverb', 'Effort', 'Traditional Wisdom'),
('الصديق وقت الضيق', 'As-sadiqu waqtu ad-diq', 'A friend in need is a friend indeed', 'Arabic Proverb', 'Friendship', 'Traditional Wisdom'),
('درهم وقاية خير من قنطار علاج', 'Dirhamu wiqayatin khayrun min qintari ''ilaj', 'An ounce of prevention is worth a pound of cure', 'Arabic Proverb', 'Prevention', 'Traditional Wisdom');

-- Update the sequence to continue from the last inserted ID
SELECT setval('quotes_id_seq', (SELECT MAX(id) FROM quotes));
