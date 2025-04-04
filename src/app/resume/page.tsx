import Tile from "../Components/tile";

import styles from './page.module.css'
import { orbitron } from "../fonts";

export default function Resume() {
    return <>
        <header className={orbitron.className}>
            <h1>
                Resume
            </h1>
            <p>
                Choose your format.
            </p>
        </header>
        <main className={styles.ResumeGrid}>
            <Tile title="PDF Version" tileStyle={styles.PDFTile} linkTo="https://collin-vesel-resumes-bucket.s3.us-east-2.amazonaws.com/Collin%20Vesel%20Resume.pdf" newTab shouldDownload/>
            <Tile title="Word Version" tileStyle={styles.WordDocTile} linkTo="https://collin-vesel-resumes-bucket.s3.us-east-2.amazonaws.com/Collin%20Vesel%20Resume.docx" newTab shouldDownload/>
        </main>
    </>
}