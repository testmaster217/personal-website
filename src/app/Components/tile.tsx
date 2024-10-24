import Link from 'next/link';

import styles from './tile.module.css';

export default function Tile({title, linkTo, x, y, width, height}: {
    title: string,
    linkTo?: string,
    x: number,
    y: number,
    width: number,
    height: number
}) {
    if (linkTo)
        return <Link
            href={linkTo}
            style={{
                gridColumn: `${x} / span ${width}`,
                gridRow: `${y} / span ${height}`
            }}
            className={styles.Tile}
        >
            <span className={styles.TileTitle}>{title}</span>
        </Link>;

    {/* This <span> is a placeholder. The actual alternate return will be
        something else, but I don;'t know what yet. */}
    return <span
        style={{
            gridColumn: `${x} / span ${width}`,
            gridRow: `${y} / span ${height}`
        }}
        className={styles.Tile}
    >
        <span className={styles.TileTitle}>{title}</span>
    </span>
}