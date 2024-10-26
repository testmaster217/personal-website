import Tile from "../Components/tile";

import styles from './page.module.css'

export default function Socials() {
    return <>
        <header>
            <h1>Socials</h1>
            <p>All the ways you can contact me.</p>
        </header>
        <main className={styles.SocialsGrid}>
            <Tile title="Email" linkTo="mailto:cvesel217@gmail.com" tileStyle={styles.EmailTile} newTab/>
            <Tile title="LinkedIn" linkTo="https://www.linkedin.com/in/collin-vesel-a547901b2/" tileStyle={styles.LinkedInTile} newTab/>
            <Tile title="GitHub" linkTo="https://github.com/testmaster217" tileStyle={styles.GitHubTile} newTab/>
        </main>
    </>
}