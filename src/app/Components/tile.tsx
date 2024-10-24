import styles from './tile.module.css';

export default function Tile({title, x, y, width, height}: {
    title: String,
    x: Number,
    y: Number,
    width: Number,
    height: Number
}) {
    return <h2 style={{
        gridColumn: `${x} / span ${width}`,
        gridRow: `${y} / span ${height}`
    }} className={styles.TileTitle}>
        {title}
    </h2>;
}