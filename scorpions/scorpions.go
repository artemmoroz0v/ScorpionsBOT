package scorpions

type Album struct {
	Description string
	Cover string
}

func SongsDataBase() map[string]Album {
	BigBase := make(map[string]Album)

	LonesomeCrow := Album {
		Description: "Album: Lonesome Crow\nYear: 1972",
		Cover: "https://e.snmc.io/i/600/s/b43c7d1fadf2dbad5be8d8031644e171/2567471/scorpions-lonesome-crow-Cover-Art.jpg",
	}
	BigBase["I'm Goin' Mad"] = LonesomeCrow
	BigBase["It All Depends"] = LonesomeCrow
	BigBase["Leave Me"] = LonesomeCrow
	BigBase["In Search Of The Peace Of Mind"] = LonesomeCrow
	BigBase["Inheritance"] = LonesomeCrow
	BigBase["Action"] = LonesomeCrow
	BigBase["Lonesome Crow"] = LonesomeCrow

	FlyToTheRainbow := Album {
		Description: "Album: Fly To The Rainbow\nYear: 1974",
		Cover: "https://i.discogs.com/TYUIZMSYoho7VZmqrOwYFsRjigQdoziCVOZo9qEVtac/rs:fit/g:sm/q:90/h:591/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTg2MzAw/MDEtMTQ4OTIyNDY0/NS03MDE5LmpwZWc.jpeg",
	}
	BigBase["Speedy's Coming"] = FlyToTheRainbow
	BigBase["They Need A Million"] = FlyToTheRainbow
	BigBase["Drifting Sun"] = FlyToTheRainbow
	BigBase["Fly People Fly"] = FlyToTheRainbow
	BigBase["This Is My Song"] = FlyToTheRainbow
	BigBase["Far Away"] = FlyToTheRainbow
	BigBase["Fly To The Rainbow"] = FlyToTheRainbow

	InTrance := Album {
		Description: "Album: In Trance\nYear: 1975",
		Cover: "https://lastfm.freetls.fastly.net/i/u/770x0/2a427f78a6dfdfcb3fedbe9586c1c56e.jpg",
	}
	BigBase["Dark Lady"] = InTrance
	BigBase["In Trance"] = InTrance
	BigBase["Life's Like A River"] = InTrance
	BigBase["Top Of The BIll"] = InTrance
	BigBase["Living And Dying"] = InTrance
	BigBase["Robot Man"] = InTrance
	BigBase["Evening Wind"] = InTrance
	BigBase["Sun In My Hand"] = InTrance
	BigBase["Longing For Fire"] = InTrance
	BigBase["NightLights"] = InTrance

	VirginKiller := Album {
		Description: "Album: Virgin Killer\nYear: 1976",
		Cover: "https://m.media-amazon.com/images/I/71NWQHr7+lL._SL1500_.jpg",
	}
	BigBase["Pictured Life"] = VirginKiller
	BigBase["Catch Your Train"] = VirginKiller
	BigBase["In Your Park"] = VirginKiller
	BigBase["Backstage Queen"] = VirginKiller
	BigBase["Virgin Killer"] = VirginKiller
	BigBase["Hell Cat"] = VirginKiller
	BigBase["Crying Days"] = VirginKiller
	BigBase["Polar Nights"] = VirginKiller
	BigBase["Yellow Raven"] = VirginKiller

	TakenByForce := Album {
		Description: "Album: Taken by Force\nYear: 1977",
		Cover: "https://upload.wikimedia.org/wikipedia/ru/1/1c/Taken_By_Force.jpg",
	}
	BigBase["Steamrock Fever"] = TakenByForce
	BigBase["We'll Burn The Sky"] = TakenByForce
	BigBase["I've Got To Be Free"] = TakenByForce
	BigBase["The Riot Of Your Time"] = TakenByForce
	BigBase["The Sails Of Charon"] = TakenByForce
	BigBase["Your Light"] = TakenByForce
	BigBase["He's A Woman - She's A Man"] = TakenByForce
	BigBase["Born To Touch Your Feelings"] = TakenByForce
	BigBase["Suspender Love"] = TakenByForce

	Lovedrive := Album {
		Description: "Album: Lovedrive\nYear: 1979",
		Cover: "https://is1-ssl.mzstatic.com/image/thumb/Music124/v4/09/ed/63/09ed635b-1fa1-fbce-cea8-7c8ea20315da/00731453478428.rgb.jpg/600x600bf-60.jpg",
	}
	BigBase["Loving You Sunday Morning"] = Lovedrive
	BigBase["Another Piece Of Meat"] = Lovedrive
	BigBase["Always Somewhere"] = Lovedrive
	BigBase["Coast To Coast"] = Lovedrive
	BigBase["Can't Get Enough"] = Lovedrive
	BigBase["Is There Anybody There"] = Lovedrive
	BigBase["Lovedrive"] = Lovedrive
	BigBase["Holiday"] = Lovedrive

	AnimalMagnetism := Album {
		Description: "Album: Animal Magnetism\nYear: 1980",
		Cover: "https://img.mvideo.ru/Pdb/40072539b.jpg",
	}
	BigBase["Make It Real"] = AnimalMagnetism
	BigBase["Don't Make No Promises"] = AnimalMagnetism
	BigBase["Hold Me Tight"] = AnimalMagnetism
	BigBase["Twentieth Century Man"] = AnimalMagnetism
	BigBase["Lady Starlight"] = AnimalMagnetism
	BigBase["Falling In Love"] = AnimalMagnetism
	BigBase["Only A Man"] = AnimalMagnetism
	BigBase["The Zoo"] = AnimalMagnetism
	BigBase["Animal Magnetism"] = AnimalMagnetism
	BigBase["Hey You"] = AnimalMagnetism

	Blackout := Album {
		Description: "Album: Blackout\nYear: 1982",
		Cover: "https://cdn1.ozone.ru/multimedia/1016706100.jpg",
	}
	BigBase["Blackout"] = Blackout
	BigBase["Can't Live Without You"] = Blackout
	BigBase["No One Like You"] = Blackout
	BigBase["You Give Me All I Need"] = Blackout
	BigBase["Now"] = Blackout
	BigBase["Dynamite"] = Blackout
	BigBase["Arizona"] = Blackout
	BigBase["China White"] = Blackout
	BigBase["When The Smoke is Going Down"] = Blackout

	LoveAtFirstSting := Album {
		Description: "Album: Love at First Sting\nYear: 1984",
		Cover: "https://lastfm.freetls.fastly.net/i/u/500x500/d76166a09b9d4ce2cce5170853dea6cb.jpg",
	}
	BigBase["Bad Boys Running Wild"] = LoveAtFirstSting
	BigBase["Rock You Like A Hurricane"] = LoveAtFirstSting
	BigBase["I'm Leaving You"] = LoveAtFirstSting
	BigBase["Coming Home"] = LoveAtFirstSting
	BigBase["The Same Thrill"] = LoveAtFirstSting
	BigBase["Big City Nights"] = LoveAtFirstSting
	BigBase["As Soon As The Good Times Roll"] = LoveAtFirstSting
	BigBase["Crossfire"] = LoveAtFirstSting
	BigBase["Still Loving You"] = LoveAtFirstSting

	SavageAmusement := Album {
		Description: "Album: Savage Amusement\nYear: 1988",
		Cover: "https://lastfm.freetls.fastly.net/i/u/ar0/e023ccdcf8665c38f9b985f6a24d0f4b.jpg",
	}
	BigBase["Don't Stop At The Top"] = SavageAmusement
	BigBase["Rhythm Of Love"] = SavageAmusement
	BigBase["Passion Rules The Game"] = SavageAmusement
	BigBase["Media Overkill"] = SavageAmusement
	BigBase["Walking On The Edge"] = SavageAmusement
	BigBase["We Let it Rock...You Let It Roll"] = SavageAmusement
	BigBase["Every Minute Every Day"] = SavageAmusement
	BigBase["Love On The Run"] = SavageAmusement
	BigBase["Believe In Love"] = SavageAmusement

	CrazyWorld := Album {
		Description: "Album: Crazy World\nYear: 1991",
		Cover: "https://is4-ssl.mzstatic.com/image/thumb/Music125/v4/ba/31/55/ba3155b3-4c11-909c-1e32-8acea6f77ac4/00042284690829.rgb.jpg/600x600bf-60.jpg",
	}
	BigBase["Tease Me Please Me"] = CrazyWorld
	BigBase["Don't Believe Her"] = CrazyWorld
	BigBase["To Be With You In Heaven"] = CrazyWorld
	BigBase["Wind Of Change"] = CrazyWorld
	BigBase["Restless Nights"] = CrazyWorld
	BigBase["Lust Or Love"] = CrazyWorld
	BigBase["Kicks After Six"] = CrazyWorld
	BigBase["Hit Between The Eyes"] = CrazyWorld
	BigBase["Money And Fame"] = CrazyWorld
	BigBase["Crazy World"] = CrazyWorld
	BigBase["Send Me An Angel"] = CrazyWorld

	FaceTheHeat := Album {
		Description: "Album: Face The Heat\nYear: 1993",
		Cover: "https://cdns-images.dzcdn.net/images/cover/4dcce5ced68c1b843423e8da91eb46c5/500x500.jpg",
	}
	BigBase["Alien Nation"] = FaceTheHeat
	BigBase["No Pain No Gain"] = FaceTheHeat
	BigBase["Someone To Touch"] = FaceTheHeat
	BigBase["Under The Same Sun"] = FaceTheHeat
	BigBase["Unholy Alliance"] = FaceTheHeat
	BigBase["Woman"] = FaceTheHeat
	BigBase["Hate To Be Nice"] = FaceTheHeat
	BigBase["Taxman Woman"] = FaceTheHeat
	BigBase["Ship Of Fools"] = FaceTheHeat
	BigBase["Nightmare Avenue"] = FaceTheHeat
	BigBase["Lonely Nights"] = FaceTheHeat
	BigBase["Destin"] = FaceTheHeat
	BigBase["Daddy's Girl"] = FaceTheHeat

	PureInstinct := Album {
		Description: "Album: Pure Instinct\nYear: 1996",
		Cover: "https://i.pinimg.com/originals/f9/27/6f/f9276fba82c5b930e51beee8e75271c4.jpg",
	}
	BigBase["Wild Child"] = PureInstinct
	BigBase["But The Best For You"] = PureInstinct
	BigBase["Does Anyone Know"] = PureInstinct
	BigBase["Stone In My Shoe"] = PureInstinct
	BigBase["Soul Behind The Face"] = PureInstinct
	BigBase["Oh Girl"] = PureInstinct
	BigBase["When You Came Into My Life"] = PureInstinct
	BigBase["Where The River Flows"] = PureInstinct
	BigBase["Time Will Call Your Name"] = PureInstinct
	BigBase["You And I"] = PureInstinct
	BigBase["Are You The One"] = PureInstinct
	BigBase["She's Knocking at My Door"] = PureInstinct

	EyeToEye := Album {
		Description: "Album: Eye II Eye\nYear: 1999",
		Cover: "https://lastfm.freetls.fastly.net/i/u/ar0/e7671c1287abeecdce3a6bb5b79a164c.jpg",
	}
	BigBase["Mysterious"] = EyeToEye
	BigBase["To Be No. 1"] = EyeToEye
	BigBase["Obsession"] = EyeToEye
	BigBase["10 Light Years Away"] = EyeToEye
	BigBase["Mind Like A Tree"] = EyeToEye
	BigBase["Eye To Eye"] = EyeToEye
	BigBase["What You Give You Get Back"] = EyeToEye
	BigBase["Skywriter"] = EyeToEye
	BigBase["Yellow Butterfly"] = EyeToEye
	BigBase["Freshly Squeezed"] = EyeToEye
	BigBase["Priscilla"] = EyeToEye
	BigBase["Du Bist So Schmutzig"] = EyeToEye
	BigBase["A Moment In A Million Years"] = EyeToEye


	Unbreakable := Album {
		Description: "Album: Unbreakable\nYear: 2004",
		Cover: "https://i.discogs.com/s0JfNzXWvENYMgmuNK8SL2WiPcZdcH05W6aZTxYN-g4/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTM4NDU1/OTAtMTQyODE3ODMx/My01Nzg3LmpwZWc.jpeg",
	}
	BigBase["New Generation"] = Unbreakable
	BigBase["Love'em Or Leave'em"] = Unbreakable
	BigBase["Deep And Dark"] = Unbreakable
	BigBase["Borderline"] = Unbreakable
	BigBase["Blood Too Hot"] = Unbreakable
	BigBase["Maybe I Maybe You"] = Unbreakable
	BigBase["Someday Is Now"] = Unbreakable
	BigBase["My City My Town"] = Unbreakable
	BigBase["Through My Eyes"] = Unbreakable
	BigBase["Can You Feel It"] = Unbreakable
	BigBase["This Time"] = Unbreakable
	BigBase["She Said"] = Unbreakable
	BigBase["Remember The Good Times"] = Unbreakable
	BigBase["Dreamers"] = Unbreakable
	BigBase["Too Far"] = Unbreakable


	HumanityHour1 := Album {
		Description: "Album: Humanity: Hour I\nYear: 2007",
		Cover: "https://upload.wikimedia.org/wikipedia/ru/9/95/Humanity_Hour_I.jpeg",
	}
	BigBase["Hour I"] = HumanityHour1	
	BigBase["The Game Of Life"] = HumanityHour1
	BigBase["We Were Born To Fly"] = HumanityHour1
	BigBase["The Future Never Dies"] = HumanityHour1
	BigBase["You're Lovin' Me To Death"] = HumanityHour1
	BigBase["321"] = HumanityHour1
	BigBase["Love Will Keep Us Alive"] = HumanityHour1
	BigBase["We Will Rise Again"] = HumanityHour1
	BigBase["Your Last Song"] = HumanityHour1
	BigBase["Love Is War"] = HumanityHour1
	BigBase["The Cross"] = HumanityHour1
	BigBase["Humanity"] = HumanityHour1
	BigBase["Cold"] = HumanityHour1


	StingInTheTail := Album {
		Description: "Album: Sting in the Tail\nYear: 2010",
		Cover: "https://lastfm.freetls.fastly.net/i/u/500x500/09d59237aed04cd8ab0b5955dac2172b.jpg",
	}
	BigBase["Raised On Rock"] = StingInTheTail
	BigBase["Sting In The Tail"] = StingInTheTail
	BigBase["Slave Me"] = StingInTheTail
	BigBase["The Good Die Young"] = StingInTheTail
	BigBase["No Limit"] = StingInTheTail
	BigBase["Rock Zone"] = StingInTheTail
	BigBase["Lorelei"] = StingInTheTail
	BigBase["Turn You On"] = StingInTheTail
	BigBase["Let's Rock"] = StingInTheTail
	BigBase["SLY"] = StingInTheTail
	BigBase["Spirit Of Rock"] = StingInTheTail
	BigBase["The Best Is Yet To Come"] = StingInTheTail

	ReturnToForever := Album {
		Description: "Album: Return to Forever\nYear: 2015",
		Cover: "https://upload.wikimedia.org/wikipedia/ru/6/6b/RTF-Cover.jpg",
	}
	BigBase["Going Out With A Bang"] = ReturnToForever
	BigBase["We Built This House"] = ReturnToForever
	BigBase["Rock My Car"] = ReturnToForever
	BigBase["House Of Cards"] = ReturnToForever
	BigBase["All For One"] = ReturnToForever
	BigBase["Rock'n'Roll Band"] = ReturnToForever
	BigBase["Catch Your Luck And Play"] = ReturnToForever
	BigBase["Rollin' Home"] = ReturnToForever
	BigBase["Hard Rockin' The Place"] = ReturnToForever
	BigBase["Eye Of The Storm"] = ReturnToForever
	BigBase["The Scratch"] = ReturnToForever
	BigBase["Gypsy Life"] = ReturnToForever
	
	RockBeliever := Album {
		Description: "Album: Rock Believer\nYear: 2022",
		Cover: "https://e.snmc.io/i/600/s/d6b9c151bfd3624f714ab459ea960972/9746111/scorpions-rock-believer-Cover-Art.jpg",
	}
	BigBase["Gas In The Tank"] = RockBeliever
	BigBase["Roots In My Boots"] = RockBeliever
	BigBase["Knock'em Dead"] = RockBeliever
	BigBase["Rock Believer"] = RockBeliever
	BigBase["Shining Of Your Soul"] = RockBeliever
	BigBase["Seventh Sun"] = RockBeliever
	BigBase["Hot And Cold"] = RockBeliever
	BigBase["When I Lay My Bones To Rest"] = RockBeliever
	BigBase["Peacemaker"] = RockBeliever
	BigBase["Call Of The WIld"] = RockBeliever
	BigBase["When You Know"] = RockBeliever
	BigBase["Shoot For Your Heart"] = RockBeliever
	BigBase["When Tomorrow Comes"] = RockBeliever
	BigBase["Unleash The Beast"] = RockBeliever
	BigBase["Crossing Borders"] = RockBeliever
	return BigBase
}
