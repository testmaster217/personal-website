import Tile from "../Components/tile";

import styles from './page.module.css'

export default function Resume() {
    return <>
        <header>
            <h1>
                Resume
            </h1>
            <p>
                Choose your format.
            </p>
        </header>
        <main className={styles.ResumeGrid}>
            <Tile title="Download PDF" tileStyle={styles.PDFTile} linkTo="./resume.pdf" newTab shouldDownload/>
            <Tile title="Download Word doc" tileStyle={styles.WordDocTile} linkTo="./resume.docx" newTab shouldDownload/>
        </main>
    </>
}