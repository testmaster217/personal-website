import styles from "./page.module.css";

import Tile from "./Components/tile";

export default function Home() {
  return <main className={styles.HomeGrid}>
    {/* TODO: Redo styling to be responsive. */}
    <Tile title="Bio" x={1} y={1} width={1} height={1}/>
    <Tile title="Resume" x={1} y={2} width={1} height={1}/>
    <Tile title="Socials" x={1} y={3} width={1} height={1}/>
    <Tile title="Featured Updates" x={2} y={1} width={3} height={3}/>
    <Tile title="All Projects" x={1} y={4} width={2} height={1}/>
    <Tile title="Toybox" x={3} y={4} width={2} height={1}/>
    <Tile title="Blog Posts" x={1} y={5} width={1} height={1}/>
    <Tile title="Thoughts" x={2} y={5} width={1} height={1}/>
    <Tile title="Not Mine, but Cool" x={3} y={5} width={1} height={1}/>
  </main>;
}
