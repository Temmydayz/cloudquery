
export const HeaderOgImage = ({ ogImage }: { ogImage?: string }) => {
    const image = ogImage || "https://cloudquery.io/og-image.jpg"

    return (
        <>
            <meta property="twitter:image" content={image} />
            <meta property="og:image" content={image} />
        </>
    )
}