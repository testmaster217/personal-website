"use client"

import Link from "next/link";

import { usePathname } from "next/navigation";
import { useEffect, useRef } from "react";
import { faHouse } from "@fortawesome/free-solid-svg-icons";

import styles from "./modalPage.module.css";
import Icon from "./icon";

export default function ModalPage({
    children
}: Readonly<{
    children: React.ReactNode
}>) {
    const pageDialog = useRef<HTMLDialogElement>(null);
    const pathname = usePathname();

    useEffect(() => {
        if (pathname.match(/\/./) && pageDialog.current && !pageDialog.current.open)
            pageDialog.current && pageDialog.current.showModal();
    }, [pathname]);

    const returnLink = pathname.substring(0, pathname.lastIndexOf("/")) === "" ? "/" : pathname.substring(0, pathname.lastIndexOf("/"));

    return <dialog ref={pageDialog} className={styles.ModalPage}>
        <div className={styles.ModalBar}>
            <Link autoFocus href={returnLink} className={styles.HomeBtn} onClick={() => pageDialog.current && pageDialog.current.close()}>
                <Icon iconName={faHouse} accessibleTitle="Home"/>
            </Link>
        </div>
        <div className={styles.ModalContent}>
            {children}
        </div>
    </dialog>
}