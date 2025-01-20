import type { Metadata } from "next";

import "./globals.css";
import styles from './homepage.module.css';
import { orbitron } from "./fonts";

import ModalPage from "./Components/modalPage";
import Tile from "./Components/tile";

export const metadata: Metadata = {
  // TODO: Add more metadata.
  title: "Collin Vesel",
  description: "My personal website. Includes some basic information about me and I hope to add (many) other features soon. :)",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        {/* I wanted open pages within the site to look like <dialog>s drawn
        on top of the homepage. This is how I did that. */}
        <ModalPage>
          {children}
        </ModalPage>

        {/* This is the homepage's content. It's here instead of in page.tsx
        because I wanted it to be in the background of other open pages and this
        is how I figued I could do it. */}
        <header>
          <h1 className={orbitron.className}>I&apos;m Collin Vesel</h1>
          <p className={orbitron.className}>and I made some things</p>
        </header>
        <main className={styles.HomeGrid}>
          <Tile title="Bio" linkTo="/bio" tileStyle={styles.BioTile}/>
          <Tile title="Resume" linkTo="/resume" tileStyle={styles.ResumeTile}/>
          <Tile title="Socials" linkTo="/socials" tileStyle={styles.SocialsTile}/>
          {/* <Tile title="Featured Updates" tileStyle={`${styles.FeaturedTile} ${styles.DisabledTile}`}/>
          <Tile title="All Projects" tileStyle={`${styles.ProjectsTile} ${styles.DisabledTile}`}/>
          <Tile title="Toybox" tileStyle={`${styles.ToyboxTile} ${styles.DisabledTile}`}/>
          <Tile title="Blog Posts" tileStyle={`${styles.BlogTile} ${styles.DisabledTile}`}/>
          <Tile title="Thoughts" tileStyle={`${styles.ThoughtsTile} ${styles.DisabledTile}`}/>
          <Tile title="Not Mine, But Cool" tileStyle={`${styles.NotMineTile} ${styles.DisabledTile}`}/> */}
        </main>
        {/* <footer className={styles.CreditNotice}>
          <small>Copyright notice for bg img goes here.</small>
        </footer> */}
      </body>
    </html>
  );
}
