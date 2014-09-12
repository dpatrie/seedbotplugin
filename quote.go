// Copyright 2014 Gabriel Guzman <gabe@lifewaza.com>
// All rights reserved.
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

// Package quote handles different type of quotes
package seedbotplugin

import (
	"github.com/seedboxtech/xmppbot"
	"math/rand"
	"strings"
	"time"
)

// Helper will implement the Helper interface
type Quote struct {
}

func (p Quote) Name() string {
	return "Quote v1.0"
}
func (p Quote) Execute(msg xmppbot.Message, bot xmppbot.Bot) error {
	if msg.From() != bot.FullName() {
		bot.Send(msg.Body())
	}
	return nil
}
// Send allows the bot to send us a message
func (h *Quote) Send(message xmppbot.Message, bot xmppbot.Bot) {
	var q = ""
	if strings.Contains(message.Body(), "code quote") {
		q = quote("code")
	}
	if strings.Contains(message.Body(), "integration quote") {
		q = quote("integration")
	}
	if strings.Contains(message.Body(), "admin quote") {
		q = quote("admin")
	}
	if len(q) > 0 {
		bot.Send(q)
	}
}

func quote(topic string) string {
	var quote string
	switch topic {
	case "code":
		quote = codeQuote()
	case "integration":
		quote = integrationQuote()
	case "admin":
		quote = adminQuote()
	default:
		quote = quoteLookup(topic)
	}
	return quote
}

// codeQuote returns an appropriate code quote
func codeQuote() string {
	quotes := []string{
		"A class name should not be a verb",
		"Leave the campground cleaner than you found it",
		"If a name requires a comment, then the name does not reveal its intent",
		"If you follow the \"one word per concept\" rule, you could end up with many classes that have, for example, an add method. As long as the parameter lists and return of the various add methods are semantically equivalent, all is well.",
		"Small! The first rule of functions is that they should be small. The second rule of functions is that they should be smaller than that!",
		"Functions should do one thing. They should to it well. They should do it only",
		"Another way to know that a function is doing more than  \"one thing\" is if you can extract another function whit a name, is not merely a restatement of its implementation",
		"In order to make sure our functions are doing \"one thing\", we need to make sure that the statements within our function are all at the same level of abstraction",
		"You know you  are working on clean code when each routine turns out to be pretty much what you expected. Half the battle to achieving that principle is choosing good names for small functions that do one thing",
		"A long descriptive name is better than a short enigmatic name. A long descriptive name is better than a long descriptive comment.",
		"The ideal number of arguments for a function is zero (niladic). Next comes one (monodic), followed closely by 2 (dyadic ). Three arguments (triadic) should be avoided where possible. More than three (polyadic) requires very special justification  .. and then shouldn't be used anyway",
		"Flag arguments are ugly. Passing a boolean into a function is a truly terrible practice. It immediately complicates the signature of the method, loudly proclaiming that this function does more than one thing. It does one thing if the flag is true and another if the flag is false!",
		"If you follow the rules herein, your functions will be short, well named, and nicely organized. But never forget that your real goal is to tell the story of the system, and the functions you write need to fit cleanly together into a clean and precise language to help you with that telling",
		"Don't comment bad code, rewrite it.  – Brian W. Kernighan and P. J. Plaugher",
		"The proper use of comments is to compensate for our failure to express ourself in code. Note that I used the word failure.",
		"Inaccurate comments are far worse than no comments at all.",
		"Few practices are as odious as commenting-out code. Like:  //$test = 1;  .. Others who see that commented-out code won't have the courage to delete it. They'll think it is there for a reason and is too important to delete. So commented-out code gathers like dregs at the bottom of a bad bottle of wine.",
		"The functionality that you create today has a good chance of changing in the next release, but the readability of your code will have a profound effect on all the changes the will ever be made.  The coding style and readability set precedents that continue to affect maintainability and extensibility long after the original code ahs been changed beyond recognition. Your style and discipline survives, even though your code does not.",
		"Objects hide their data behind abstractions and expose functions that operate on that data. Data structure expose their data and have no meaningful functions.",
		"Read the book clean code. We have some at the office. Ask Pascal",
		"We don't usually expect information to be going out through the arguments",
		"Clear and expressive code with few comments is far superior to cluttered and complex code with lots of comments",
		"Rather than spend your time writing the comments that explain the mess you've made, spend it cleaning that mess",
		"You will be challenged to think about what's right about that code and what's wrong with it",
		"Create informative error messages and pass them along with your exceptions. Mention the operation that failed and the type of failure. If you are loggin in your application, pass along enough information to be able to log the error in your catch",
		"Clean code is readable, but it must also be robust. These are not conflicting goals. We can write robust clean code if we see error handling as a separate concern, something that is viewable independently of our main logic. To the degree that we are able to do that, we can reason about it independently and we can make great strides in the maintainability of our code",
		"3 laws of TDD.  1st: You may not write production code until you have written a failing unit test.2nd: You may not write more of a unit test than is sufficient to fail and not compiling is failing. 3rd: You may not write more production code than is sufficient to pass the currently failing test",
		"One difference between a smart programmer and a professional programmer is that the professional understands that clarity is king.",
		"There are two ways of constructing a software design: One way is to make it so simple that there are obviously no deficiencies and the other way is to make it so complicated that there are no obvious deficiencies.  — C.A.R. Hoare, The 1980 ACM Turing Award Lecture",
		"The computing scientist's main challenge is not to get confused by the complexities of his own making. — E. W. Dijkstra",
		"Always code as if the person who ends up maintaining your code is a violent psychopath who knows where you live.",
		"Any 3rd party system\n that I have to integrate\n with, was written by\n a drunken monkey\n typing with his feet",
		"Any fool can write code that a computer can understand. Good programmers write code that humans can understand.",
		"Singletons are the path of the dark side. Singletons lead to implicit dependencies. Implicit dependencies lead to global state. Global state leads to suffering.",
		"'First, catch the rabbit.'\n - a recipe for rabbit stew",
		"It is often a mistake to make a priori judgments about what parts of a program are really critical, since the universal experience of programmers who have been using measurement tools has been that their intuitive guesses fail.\n - Donald Knuth",
		"As I've said, much of this advice --in particular, the advice to write a good clean program first and optimize it later-- is well worn... For everyone who finds nothing new in this column, there exists another challenge: how to make it so there is no need to rewrite it in 10 years.\n -Martin Fowler (10 years ago), Yet Another Optimization Article",
		"I don't believe in the Performance Fairy.\n -Jeff Rothschild",
		"Prudent physicists --those who want to avoid false leads and dead ends-- operate according to a long standing principle: Never start a lengthy calculation until you know the range of values within which the answer is likey to fall (and, equally important,  the range within the answer is unlikely to fall).\n -Hans Christian van Baeyer, The Fermi Solution",
		"It is better to do the right problem the wrong way than the wrong problem the right way.\n -Richard Hamming, The Art of Doing Science & Engineering",
		"One of my most productive days was throwing away 1000 lines of code.",
		"Deleted code is debugged code. — Jeff Sickel",
		"Debugging is twice as hard as writing the code in the first place. Therefore, if you write the code as cleverly as possible, you are, by definition, not smart enough to debug it. — Brian W. Kernighan and P. J. Plauger in The Elements of Programming Style.",
		"Controlling complexity is the essence of computer programming. — Brian Kernighan",
		"Simplicity is prerequisite for reliability.  — Edsger W. Dijkstra",
		"Essentially everyone, when they first build a distributed application, makes the following eight assumptions. All prove to be false in the long run and all cause big trouble and painful learning experiences.\n The network is reliable\n Latency is zero\n Bandwidth is infinite\n The network is secure\n Topology doesn't change\n There is one administrator\n Transport cost is zero\n The network is homogeneous\n — Peter Deutsch",
		"First, solve the problem. Then, write the code.        — John Johnson",
		"Complexity kills. It sucks the life out of developers, it makes products difficult to plan, build and test, it introduces security challenges and it causes end-user and administrator frustration.   — Ray Ozzie",
		"Any program that tries to be so generalized and configurable that it could handle any kind of task will either fall short of this goal, or will be horribly broken. — Chris Wenham",
		"Debugging time increases as a square of the program's size. — Chris Wenham",
		"Before software can be reusable it first has to be usable.",
		"Code never lies, comments sometimes do.   — Ron Jeffries",
		"... the cost of adding a feature isn't just the time it takes to code it. The cost also includes the addition of an obstacle to future expansion. ... The trick is to pick the features that don't fight each other.        — John Carmack",
		"Software is like entropy. It is difficult to grasp, weighs nothing, and obeys the second law of thermodynamics; i.e. it always increases.",
		"Software gets slower faster than hardware gets faster.",
		"So much complexity in software comes from trying to make one thing do two things. — Ryan Singer",
		"The simplest implementation is almost always one that has already been implemented and proven scalable.",
		"There are other factors that can contribute to software rot [...] but neglect accelerates the rot faster than any other factor.",
		"functions should either do something or answer something, but not both",
		"Care about your code. If you don’t, then chances are no-one else will either.",
	}
	rand.Seed(time.Now().UnixNano())
	quote := quotes[rand.Intn(len(quotes))]
	return quote
}

func integrationQuote() string {
	quotes := []string{
		"Browsers have the bad habit to set some default CSS to many of the tags that are regularly used. Many different browser use different default CSS, which might make the sites look different from one browser to another. Use a RESET css from http://www.meyerweb.com to prevent that.",
		"Semantic HTML is when you favor meaning over presentation in your HTML markup. It helps define the role of the different parts of content within the page. Using DIVS everywhere does not tell much of the role this tag is playing.",
		"If you're doing a list of links, please create a list! <ul><li>, it's that simple",
		"Another face of semantics is when using CSS classes. It is very important that your classes shows what role the element has, and not what it looks like. A class called \"BoldRed\" will lost all its sense if you change the color to blue. You should consider using a classname like \"ImportantText\"",
		"Keep your CSS short by using shorthand syntax. Examples : font: bold 12px/20px Arial, Verdana; margin: 10px 20px 5px; border: 1px solid #000; background: url(image.jpg) 0 0 no-repeat;",
		"Sprites is when you combine multiple images into a big one in order to reduce the number of HTTP_REQUESTS, thus also reducing loading time. The smallest amount of different images you load, the less time it takes to load the entire page. By combining multiple images into one, you get a bigger images, but it it still profitable than having 20 small images loading in parallel.",
		"Did you know? Organizing your sprites horizontally as opposed to vertically usually results in smaller file sizes.",
		"Did you know? Combining similar colors in a sprite helps you keep the color count low, ideally under 256 colors so to fit in a PNG8.",
		"\"Be mobile-friendly\" and don't leave big gaps between the images in a sprite. This doesn't affect the file size as much but requires less memory for the user agent to decompress the image into a pixel map. 100x100 image is 10 thousand pixels, where 1000x1000 is 1 million pixels",
		"Did you know? Optimizing your PNG images with http://www.tinypng.org can save an enormous amount of weight without losing quality.",
		"You can save a tremendous amount page weight if you do your gradients with CSS3. Try colorzilla! http://www.colorzilla.com/gradient-editor/",
		"You should use JPEGs for a photo, GIFs for fla-color images like logos and buttons, and PNGs for images that require more than 256 colors.",
		"To reduce page load time, you should : combine images in sprites, never use images to display text, group CSS selectors having similar properties, avoid using CSS expressions and filters, keep your stylesheets in the <head> of the page.",
		"Doing sprites whenever you can should become a reflex. http://cdn.memegenerator.net/instances/400x/20584328.jpg",
	}
	rand.Seed(time.Now().UnixNano())
	quote := quotes[rand.Intn(len(quotes))]
	return quote
}

func adminQuote() string {
	quotes := []string{
		"The problem with troubleshooting is that trouble shoots back.  ~Author Unknown",
		"Computers have lots of memory but no imagination.  ~Author Unknown",
		"There are three kinds of death in this world.  There's heart death, there's brain death, and there's being off the network.  ~Guy Almes",
		"RAM disk is not an installation procedure.  ~Author Unknown",
		"Those parts of the system that you can hit with a hammer (not advised) are called hardware; those program instructions that you can only curse at are called software.  ~Author Unknown",
		"If a trainstation is where the train stops, what's a workstation?  ~Author Unknown",
		"Jesus saves!  The rest of us better make backups.  ~Author Unknown",
		"In God we trust, all others we virus scan.  ~Author Unknown",
		"The question of whether computers can think is just like the question of whether submarines can swim.  ~Edsger W. Dijkstra",
		"Why did the sysadmin cross the road?  To get coffee, why else would one be outside?  ~Author Unknown",
		"To err is human, to really foul things up requires a computer.  ~Bill Vaughan, 1969 (Thanks, Garson O'Toole!)",
		"There are two major products that came out of Berkeley:  LSD and UNIX.  We do not believe this to be a coincidence.  ~Jeremy S. Anderson",
	}
	rand.Seed(time.Now().UnixNano())
	quote := quotes[rand.Intn(len(quotes))]
	return quote
}

func quoteLookup(string) string {
	quote := "not yet implemented"
	return quote
}
