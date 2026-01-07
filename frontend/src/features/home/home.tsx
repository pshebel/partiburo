import { getHome } from '../../hooks/home';
import { getGuest } from '../../hooks/identity'
import { Link } from 'react-router-dom'
export const Home = () => {
    const { data, isLoading, error } = getHome();
    const guest_id = getGuest()
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
                <div>Date: {data.Date}</div>
                <div>Time: {data.Time}</div>
                <div>Address: {data.Address}</div>
            </div>
            <div>
                <h1>announcements</h1>
                {data.Announcements.sort((a,b) => Date.parse(b.CreatedAt) - Date.parse(a.CreatedAt)).map((a, i) => (
                    <div key={i}>
                        <h3>{a.Header}</h3>
                        <div>{a.Body}</div>
                    </div>
                ))}
            </div>
            <div>
                <h1>guests</h1>
                {data.Guests.sort((a,b) => Date.parse(b.createdAt) - Date.parse(a.createdAt)).map((g, i) => {
                    return (
                    <div key={i}>
                        <span>{g.name}</span>
                        <span>{g.status}</span>
                        {g.id === guest_id && <Link to="/guest">Edit</Link>}
                    </div>
                    )}
                )}
            </div>
            <div>
                <h1>posts</h1>
                <h2><Link to="/post">create post</Link></h2>
                {data.Posts.sort((a,b) => Date.parse(b.createdAt) - Date.parse(a.createdAt)).map((p, i) => (
                    <div key={i}>
                        <span>{p.name}</span>
                        <span>{p.body}</span>
                    </div>
                ))}
            </div>
        </div>
    )
}