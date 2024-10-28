// WARNING: This file causes errors when trying to make a production build.
// I have no idea how to fix them other than removing the file and doing
// something different. I hope to figure this out at some point.

"use client"

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { IconDefinition } from "@fortawesome/free-solid-svg-icons";
import { useEffect, useState } from "react";

export default function Icon({iconName, accessibleTitle}: {
    iconName: IconDefinition,
    accessibleTitle: string
}) {
    // Use useEffect and a state variable to check if the component is on the client.
    // This prevents hydration errors caused by randomized parts of the FA icon's aria label.
    const [isClient, setIsClient] = useState(false);
    useEffect(() => {
        setIsClient(true);
    }, []);

    return isClient ? <FontAwesomeIcon icon={iconName} title={accessibleTitle}/> : <></>;
}