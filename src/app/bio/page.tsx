import Link from "next/link";

export default function Bio() {
    return <>
        <header>
            <h1>Bio</h1>
            <p>A little bit about me.</p>
        </header>
        <main>
            <article>
                {/* TODO: Add photo of me. */}
                <p>My name is Collin Vesel, and I am a software engineer.</p>
                <p>
                    I fell in love with programming and the creative power it gave 
                    me in eighth grade, and have been building up my skills and my 
                    ideas for what to do with them ever since. Most of my experience 
                    so far has been in web, desktop, and game development – 
                    specifically with Java, C#, and JavaScript mostly – but I am more 
                    than willing to learn any tools or technologies that I need to.
                </p>
                <p>
                    I also have high standards for the work that I do and am very 
                    detail-oriented. In college, I would often start working on 
                    major assignments as soion as they were available or anounced, 
                    continue working on it until right before the deadline, think, 
                    "It's not as good as I'd like, but it's good enough" as I was 
                    about to submit it, and then get full credit for it because I 
                    had done a better job than I realized or was trying to do.
                </p>
                <p>
                    All of this led me to complete a software engineering internship 
                    at Enterprise Holdings in the summer of 2022, graduate summa cum 
                    laude with a Bachelor's of Science in Computer Science the 
                    following spring, earn a Coursera certification in frontend 
                    development from Meta in the fall of 2024, and develop multiple 
                    personal projects like this website. (You can check out several 
                    projects of mine that I thought were worth sharing with the world 
                    on <Link href="https://github.com/testmaster217">my GitHub profile</Link>.)
                </p>
                <p>
                    As college wrapped up and after I graduated, I translated my high 
                    standards and desire to do the best work that I can into an 
                    understanding that software engineering is a <i>lot</i> more than 
                    just programming. It also involves knowing what problems to solve 
                    (by listening to stakeholders) and how best to solve them. It 
                    involves working with team members and effectively managing the 
                    project. It involves knowing when to optimize and when there are 
                    more important things to worry about (and if it is time to optimize, 
                    it involves knowing what needs the optimization).
                </p>
                <p>
                    I have the tools – or the willingness to learn. I have some of the 
                    wisdom – and an intent to get wiser. I have an eye for details and 
                    am building one for design – even if my vision impairment keeps me 
                    from seeing much else without Windows Magnifier (which reminds me, 
                    I should probably use that to finsih writing this...🙃). And I have 
                    a <i>massive</i> amount of ideas for what to do with all of this 
                    (which I affectionately refer to as my "Dragon's Hoard of iCloud Notes").
                </p>
                <p>
                    The only thing I don't have is 3+ years of professional experience.
                </p>
                <p>
                    If you would like to help me build up that experience, you can email 
                    me or message me on LinkedIn by going to <Link href="/socials">the 
                    "Socials" page</Link> on this site. I will respond to you quickly.
                </p>
                <p>
                    I look forward to hearing from you. :)
                </p>
                <p>
                    P.S., I know there's not much to this site yet. It's what I thought 
                    would be a minimum viable product right now. However, I have many 
                    updates planned and I'm excited for you to see what else I add. ;)
                </p>
            </article>
        </main>
    </>
}