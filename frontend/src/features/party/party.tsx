import { getParty } from '../../hooks/party';


export const Party = () => {
    const { data, isLoading, error } = getParty();
    if (isLoading) {
        return (
            <div>
                <span>Loading</span>
            </div>
        )
    }

    if (error) {
        return (
            <div>
                <span>Error: {error.message}</span>
            </div>
        )
    }

    if (data === undefined) {
        return (
            <div>
                <span>Error: Failed to load data</span>
            </div>
        )
    }
    return (
        <div>
            <div>
                <h1>about</h1>
                <h2>{data.Title}</h2>
                <div>{data.Description}</div>
            </div>
            <div>
                <h1>announcements</h1>
                {data.Announcements.sort((a,b) => Date.parse(a.CreatedAt) - Date.parse(b.CreatedAt)).map((a, i) => (
                    <div key={i}>
                        <h3>{a.Header}</h3>
                        <div>{a.Body}</div>
                    </div>
                ))}
            </div>
            <div>
                <h1>guests</h1>
                {data.Guests.sort((a,b) => Date.parse(a.CreatedAt) - Date.parse(b.CreatedAt)).map((g, i) => (
                    <div key={i}>
                        <span>{g.Name}</span>
                        <span>{g.Status}</span>
                    </div>
                ))}
            </div>
            <div>
                <h1>posts</h1>
                {data.Posts.sort((a,b) => Date.parse(a.CreatedAt) - Date.parse(b.CreatedAt)).map((p, i) => (
                    <div key={i}>
                        <span>{p.Name}</span>
                        <span>{p.Body}</span>
                    </div>
                ))}
            </div>
        </div>
    )
}