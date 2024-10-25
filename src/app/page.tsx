import styles from "./page.module.css";

import Tile from "./Components/tile";

export default function Home() {
  return <>
    <header className={styles.SiteTitle}>
      <h1 className={styles.ClearMargin}>I'm Collin Vesel</h1>
      <p>and I made some things</p>
    </header>
    <main className={styles.HomeGrid}>
      {/* TODO: Redo styling to be responsive. */}
      <Tile title="Bio" linkTo="/bio" x={1} y={1} width={1} height={1}/>
      <Tile title="Resume" linkTo="/resume" x={2} y={1} width={2} height={1}/>
      <Tile title="Socials" linkTo="/socials" x={1} y={2} width={2} height={1}/>
      <Tile title="Featured Updates" x={1} y={3} width={3} height={3}/>
      <Tile title="All Projects" x={1} y={6} width={1} height={1}/>
      <Tile title="Toybox" x={2} y={6} width={1} height={1}/>
      <Tile title="Blog Posts" x={1} y={7} width={1} height={1}/>
      <Tile title="Thoughts" x={2} y={7} width={1} height={1}/>
      <Tile title="Not Mine, but Cool" x={3} y={6} width={1} height={2}/>
    </main>
    {/* <footer>
      <small>Copyright notice for bg img goes here.</small>
    </footer> */}
  </>;
}
