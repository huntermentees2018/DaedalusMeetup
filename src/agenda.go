package src

import (
	"math/rand"
	"time"
)

func GetAgenda() string {
	question := grabRandomQuestion()
	return buildAgenda(question)
}

func buildAgenda(question string) string {
	return "Hello friends! Ready to get your 1:1 on? Here’s the agenda for the best meeting of your life:\n" +
		"I. Introductions: What's your major, what's your favorite class? least favorite?\n" +
		"II. " + question + "\n" +
		"III. Talk amongst yourselves. What is something about yourself that you don't think the other person knows?\n" +
		"\n" +
		"Happy 1:1ing!\n"
}

func grabRandomQuestion() string {
	rand.Seed(time.Now().UnixNano())

	questions := rawAgendaQuestions()
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	return questions[0]
}

func rawAgendaQuestions() []string {
	return []string{
		"Would you rather always be 10 minutes late or always be 20 minutes early?",
		"Would you rather lose all of your money and valuables or all of the pictures you have ever taken?",
		"Would you rather be able to see 10 minutes into your own future or 10 minutes into the future of anyone but yourself?",
		"Would you rather be famous when you are alive and forgotten when you die or unknown when you are alive but famous after you die?",
		"Would you rather go to jail for 4 years for something you didn’t do or get away with something horrible you did but always live in fear of being caught?",
		"Would you rather your shirts be always two sizes too big or one size too small?",
		"Would you rather live your entire life in a virtual reality where all your wishes are granted or in the real world?",
		"Would you rather be alone for the rest of your life or always be surrounded by annoying people?",
		"Would you rather never use social media sites / apps again or never watch another movie or TV show?",
		"Would you rather be the first person to explore a planet or be the inventor of a drug that cures a deadly disease?",
		"Would you rather have a horrible short term memory or a horrible long term memory?",
		"Would you rather be completely invisible for one day or be able to fly for one day?",
		"Would you rather be locked in a room that is constantly dark for a week or a room that is constantly bright for a week?",
		"Would you rather live without the internet or live without AC and heating?",
		"Would you rather find your true love or a suitcase with five million dollars inside?",
		"Would you rather be able to teleport anywhere or be able to read minds?",
		"Would you rather be feared by all or loved by all?",
		"Would you rather be transported permanently 500 years into the future or 500 years into the past?",
		"Would you rather never be able to use a touchscreen or never be able to use a keyboard and mouse?",
		"Would you rather have everything you eat be too salty or not salty enough no matter how much salt you add?",
		"Would you rather have hands that kept growing as you got older or feet that kept growing as you got older?",
		"Would you rather be unable to use search engines or unable to use social media?",
		"Would you rather give up bathing for a month or give up the internet for a month?",
		"Would you rather go back to age 5 with everything you know now or know now everything your future self will learn?",
		"Would you rather relive the same day for 365 days or lose a year of your life?",
		"Would you rather be able to control animals (but not humans) with your mind or control electronics with your mind?",
		"Would you rather never have to work again or never have to sleep again (you won’t feel tired or suffer negative health effects)?",
		"Would you rather get one free round trip international plane ticket every year or be able to fly domestic anytime for free?",
		"Would you rather be able to be free from junk mail or free from email spam for the rest of your life?",
		"Would you rather have an unlimited international first class ticket or never have to pay for food at restaurants?",
		"Would you rather see what was behind every closed door or be able to guess the combination of every safe on the first try?",
		"Would you rather live in virtual reality where you are all powerful or live in the real world and be able to go anywhere but not be able to interact with anyone or anything?",
		"Would you rather never be able to eat meat or never be able to eat vegetables?",
		"Would you rather give up watching TV / movies for a year or give up playing games for a year?",
		"Would you rather always be able to see 5 minutes into the future or always be able to see 100 years into the future?",
		"Would you rather super sensitive taste or super sensitive hearing?",
		"Would you rather be a practicing doctor or a medical researcher?",
		"Would you rather be married to a 10 with a bad personality or a 6 with an amazing personality?",
		"Would you rather never be able to drink sodas like coke again or only be able to drink sodas and nothing else?",
		"Would you rather have amazingly fast typing / texting speed or be able to read ridiculously fast?",
		"Would you rather know the history of every object you touched or be able to talk to animals?",
		"Would you rather be a reverse centaur or a reverse mermaid/merman?",
		"Would you rather have constantly dry eyes or a constant runny nose?",
		"Would you rather be a famous director or a famous actor?",
		"Would you rather not be able to open any closed doors (locked or unlocked) or not be able to close any open doors?",
		"Would you rather give up all drinks except for water or give up eating anything that was cooked in an oven?",
		"Would you rather be constantly tired no matter how much you sleep or constantly hungry no matter what you eat? Assuming that there are no health problems besides the feeling of hunger and sleepiness.",
		"Would you rather have to read aloud every word you read or sing everything you say out loud?",
		"Would you rather have whatever you are thinking appear above your head for everyone to see or have absolutely everything you do live streamed for anyone to see?",
		"Would you rather be put in a maximum security federal prison with the hardest of the hardened criminals for one year or be put in a relatively relaxed prison where wall street types are held for ten years?",
		"Would you rather have a clown only you can see that follows you everywhere and just stands silently in a corner watching you without doing or saying anything or have a real life stalker who dresses like the Easter bunny that everyone can see?",
		"Would you rather have a completely automated home or a self-driving car?",
		"Would you rather work very hard at a rewarding job or hardly have to work at a job that isn’t rewarding?",
		"Would you rather be held in high regard by your parents or your friends?",
		"Would you rather be an amazing painter or a brilliant mathematician?",
		"Would you rather be reincarnated as a fly or just cease to exist after you die?",
		"Would you rather be able to go to any theme park in the world for free for the rest of your life or eat for free at any drive through restaurant for the rest of your life?",
		"Would you rather be only able to watch the few movies with a rotten tomatoes score of 95-100% or only be able to watch the majority of movies with a rotten tomatoes score of 94% and lower?",
		"Would you rather never lose your phone again or never lose your keys again?",
		"Would you rather have one real get out of jail free card or a key that opens any door?",
		"Would you rather have a criminal justice system that actually works and is fair or an administrative government that is free of corruption?",
		"Would you rather have real political power but be relatively poor or be ridiculously rich and have no political power?",
		"Would you rather have the power to gently nudge anyone’s decisions or have complete puppet master control of five people?",
		"Would you rather have everyone laugh at your jokes but not find anyone else’s jokes funny or have no one laugh at your jokes but you still find other people’s jokes funny?",
		"Would you rather be the absolute best at something that no one takes seriously or be well above average but not anywhere near the best at something well respected?",
		"Would you rather lose the ability to read or lose the ability to speak?",
		"Would you rather live under a sky with no stars at night or live under a sky with no clouds during the day?",
		"Would you rather humans go to the moon again or go to mars?",
		"Would you rather never get angry or never be envious?",
		"Would you rather have free Wi-Fi wherever you go or be able to drink unlimited free coffee at any coffee shop?",
		"Would you rather be compelled to high five everyone you meet or be compelled to give wedgies to anyone in a green shirt?",
		"Would you rather live in a house with see-through walls in a city or in the same see-though house but in the middle of a forest far from civilization?",
		"Would you rather take amazing selfies but all of your other pictures are horrible or take breathtaking photographs of anything but yourself?",
		"Would you rather use a push lawn mower with a bar that is far too high or far too low?",
		"Would you rather be able to dodge anything no matter how fast it’s moving or be able ask any three questions and have them answered accurately?",
		"Would you rather live on the beach or in a cabin in the woods?",
		"Would you rather lose your left hand or right foot?",
		"Would you rather face your fears or forget that you have them?",
		"Would you rather be forced to dance every time you heard music or be forced to sing along to any song you heard?",
		"Would you rather have skin that changes color based on your emotions or tattoos appear all over your body depicting what you did yesterday?",
		"Would you rather live in a utopia as a normal person or in a dystopia but you are the supreme ruler?",
		"Would you rather snitch on your best friend for a crime they committed or go to jail for the crime they committed?",
		"Would you rather have everything on your phone right now (browsing history, photos, etc.) made public to anyone who Google’s you name or never use a cell phone again?",
		"Would you rather eat a box of dry spaghetti noodles or a cup of uncooked rice?",
		"Would you rather wake up as a new random person every year and have full control of them for the whole year or once a week spend a day inside a stranger without having any control of them?",
		"Would you rather be born again in a totally different life or born again with all the knowledge you have now?",
		"Would you rather be lost in a bad part of town or lost in the forest?",
		"Would you rather never get a paper cut again or never get something stuck in your eye again?",
		"Would you rather randomly time travel +/- 20 years every time you fart or teleport to a different place on earth (on land, not water) every time you sneeze?",
		"Would you rather the aliens that make first contact be robotic or organic?",
		"Would you rather be famous but ridiculed or be just a normal person?",
		"Would you rather be an amazing virtuoso at any instrument but only if you play naked or be able to speak any language but only if close your eyes and dance while you are doing it?",
		"Would you rather have a flying carpet or a car that can drive underwater?",
		"Would you rather be an amazing artist but not be able to see any of the art you created or be an amazing musician but not be able to hear any of the music you create?",
		"Would you rather there be a perpetual water balloon war going on in your city / town or a perpetual food fight?",
		"Would you rather find five dollars on the ground or find all of your missing socks?",
		"Would you rather never have another embarrassing fall in public or never feel the need to pass gas in public again?",
		"Would you rather be able to talk to land animals, animals that fly, or animals that live under the water?",
		"Would you rather lose your best friend or all of your friends except for your best friend?",
		"Would you rather it be impossible for you to be woken up for 11 straight hours every day but you wake up feeling amazing or you can be woken up normally but never feel totally rested?",
		"Would you rather everything you dream each night come true when you wake up or everything a randomly chosen person dreams each night come true when they wake up?",
		"Would you rather have every cat or dog that gets lost end up at your house or everyone’s clothes that they forget in the dryer get teleported to your house?",
		"Would you rather never be stuck in traffic again or never get another cold?",
		"Would you rather know how above or below average you are at everything or know how above or below average people are at one skill / talent just by looking at them?",
		"Would you rather have a cute well behaved child that stays at an age of your choosing for their entire life or a child that develops from a baby to 18 years old in 2 years and then ages normally?",
		"Would you rather it never stops snowing (the snow never piles up) or never stops raining (the rain never causes floods)?",
		"Would you rather have a bottomless box of Legos or a bottomless gas tank?",
		"Would you rather wake up each morning to find that a random animal appendage has replaced your non dominant arm or permanently replace your bottom half with an animal bottom of your choice?",
		"Would you rather fight for a cause you believe in but doubt will succeed or fight for a cause that you only partially believe in but have a high chance of your cause succeeding?",
		"Would you rather have no fingers or no elbows?",
		"Would you rather have edible spaghetti hair that regrows every night or sweat maple syrup?",
		"Would you rather have everything in your house perfectly organized by a professional or have a professional event company throw the best party you’ve ever been to, in your honor?",
		"Would you rather have all traffic lights you approach be green or never have to stand in line again?",
		"Would you rather have all of your clothes fit perfectly or have the most comfortable pillow, blankets, and sheets in existence?",
		"Would you rather wake up in the middle of an unknown desert or wake up in a row boat on an unknown body of water?",
		"Would you rather be famous for inventing a deadly new weapon or invent something that helps the world but someone else gets all the credit for inventing it?",
		"Would you rather 5% of the population have telepathy or 5% of the population have telekinesis? You are not part of the 5% that has telepathy or telekinesis.",
		"Would you rather earbuds and headphones never sit right on / in your ears or have all music either slightly too quiet or slightly too loud?",
		"Would you rather become twice as strong when both of your fingers are stuck in your ears or crawl twice as fast as you can run?",
		"Would you rather live in a giant desert or a giant dessert?",
		"Would you rather have everything you draw become real but be terrible at drawing or be able to fly but only as fast as you can walk?",
		"Would you rather have a map that shows you the location of anything you want to find and can be used again and again but has a margin of error of up to a mile or a device that allows you to find the location of anything you want with incredible accuracy but can only be used three times?",
		"Would you rather never run out of battery power for whatever phone and tablet you own or always have free Wi-Fi wherever you go?",
	}
}
