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
            <Tile title="Download PDF" tileStyle={styles.PDFTile} linkTo="/Resumes/Collin Vesel Resume.pdf" newTab shouldDownload/>
            <Tile title="Download Word doc" tileStyle={styles.WordDocTile} linkTo="/Resumes/Collin Vesel Resume.docx" newTab shouldDownload/>
        </main>
    </>
}