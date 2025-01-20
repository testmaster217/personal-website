"use client"

import Link from "next/link";

import { usePathname } from "next/navigation";
import { useEffect, useRef } from "react";

import styles from "./modalPage.module.css";
import { michroma } from "../fonts";

export default function ModalPage({
    children
}: Readonly<{
    children: React.ReactNode
}>) {
    const pageDialog = useRef<HTMLDialogElement>(null);
    const pathname = usePathname();

    useEffect(() => {
        if (pathname.match(/\/./) && pageDialog.current && !pageDialog.current.open)
            pageDialog.current.showModal();
    }, [pathname]);

    const returnLink = pathname.substring(0, pathname.lastIndexOf("/")) === "" ? "/" : pathname.substring(0, pathname.lastIndexOf("/"));

    return <dialog ref={pageDialog} className={styles.ModalPage}>
        <div className={styles.ModalBar}>
            <Link autoFocus href={returnLink} className={`${styles.HomeBtn} ${michroma.className}`} onClick={() => pageDialog.current && pageDialog.current.close()}>
                HOME
            </Link>
        </div>
        <div className={styles.ModalContent}>
            {children}
        </div>
    </dialog>
}