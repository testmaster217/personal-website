import styles from "./page.module.css";

import Tile from "./Components/tile";

export default function Home() {
  return <>
    <header className={styles.SiteTitle}>
      <h1 className={styles.ClearMargin}>I'm Collin Vesel</h1>
      <p>and I made some things</p>
    </header>
    <main className={styles.HomeGrid}>
      <Tile title="Bio" linkTo="/bio" tileStyle={styles.BioTile}/>
      <Tile title="Resume" linkTo="/resume" tileStyle={styles.ResumeTile}/>
      <Tile title="Socials" linkTo="/socials" tileStyle={styles.SocialsTile}/>
      <Tile title="Featured Updates" tileStyle={styles.FeaturedTile}/>
      <Tile title="All Projects" tileStyle={styles.ProjectsTile}/>
      <Tile title="Toybox" tileStyle={styles.ToyboxTile}/>
      <Tile title="Blog Posts" tileStyle={styles.BlogTile}/>
      <Tile title="Thoughts" tileStyle={styles.ThoughtsTile}/>
      <Tile title="Not Mine, but Cool" tileStyle={styles.NotMineTile}/>
    </main>
    {/* <footer>
      <small>Copyright notice for bg img goes here.</small>
    </footer> */}
  </>;
}
