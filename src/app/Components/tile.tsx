import Link from 'next/link';

import styles from './tile.module.css';

export default function Tile({title, tileStyle, linkTo, newTab, shouldDownload}: {
    title: string,
    tileStyle: string,
    linkTo?: string,
    newTab?: boolean,
    shouldDownload?: boolean
}) {
    if (linkTo)
        return <Link href={linkTo} className={`${styles.Tile} ${tileStyle}`} target={newTab ? "_blank" : "_self"} download={shouldDownload}>
            <span className={styles.TileTitle}>{title}</span>
        </Link>;

    {/* This <span> is a placeholder. The actual alternate return will be
        something else, but I don;'t know what yet. */}
    return <span className={`${styles.Tile} ${tileStyle}`}>
        <span className={styles.TileTitle}>{title}</span>
    </span>
}