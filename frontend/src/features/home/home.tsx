import { getHome } from '../../hooks/home';
import { getGuest } from '../../hooks/identity';
import { Link } from 'react-router-dom';

export const Home = () => {
    const { data, isLoading, error } = getHome();
    const guest_id = getGuest();

    // Loading State
    if (isLoading) {
        return (
            <div className="flex items-center justify-center min-h-screen bg-gray-50">
                <span className="text-lg font-medium animate-pulse text-gray-600">Loading...</span>
            </div>
        );
    }

    // Error State
    if (error || !data) {
        return (
            <div className="p-6 m-4 bg-red-50 border border-red-200 rounded-lg text-red-700">
                <span className="font-bold">Error:</span> {error?.message || "Failed to load data"}
            </div>
        );
    }

    return (
        <div className="max-w-4xl mx-auto p-6 space-y-12 bg-white min-h-screen">
            
            {/* 1. About Section */}
            <section className="border-b pb-8">
                <h1 className="text-xs font-bold uppercase tracking-widest text-blue-600 mb-2">About</h1>
                <h2 className="text-4xl font-extrabold text-gray-900 mb-4">{data.Title}</h2>
                <p className="text-lg text-gray-600 leading-relaxed mb-6 whitespace-pre-line">
                    {data.Description}
                </p>
                
                <div className="grid grid-cols-1 md:grid-cols-3 gap-4 text-sm font-medium text-gray-500">
                    <div className="flex items-center gap-2">📅 {data.Date}</div>
                    <div className="flex items-center gap-2">⏰ {data.Time}</div>
                    <div className="flex items-center gap-2">📍 {data.Address}</div>
                </div>
            </section>

            {/* 2. Announcements */}
            <section>
                <h1 className="text-xs font-bold uppercase tracking-widest text-orange-600 mb-6">Announcements</h1>
                <div className="space-y-4">
                    {data.Announcements.sort((a,b) => Date.parse(b.CreatedAt) - Date.parse(a.CreatedAt)).map((a, i) => (
                        <div key={i} className="p-4 bg-orange-50 rounded-xl border border-orange-100">
                            <h3 className="font-bold text-gray-900">{a.Header}</h3>
                            <div className="text-gray-700 mt-1">{a.Body}</div>
                        </div>
                    ))}
                </div>
            </section>

            {/* 3. Guests List */}
            <section>
                <h1 className="text-xs font-bold uppercase tracking-widest text-green-600 mb-6">Guest List ({data.Going} going)</h1>
                <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
                    {data.Guests.sort((a,b) => Date.parse(b.createdAt) - Date.parse(a.createdAt)).map((g, i) => (
                        <div key={i} className="flex items-center justify-between p-3 bg-gray-50 rounded-lg border">
                            <div className="flex flex-col">
                                <span className="font-semibold text-gray-800">{g.name}</span>
                                <span className="text-xs text-gray-500 italic">{g.status}</span>
                            </div>
                            {g.id === guest_id && (
                                <Link to="/guest" className="text-xs bg-white border px-3 py-1 rounded hover:bg-gray-100 transition shadow-sm">
                                    Edit Profile
                                </Link>
                            )}
                        </div>
                    ))}
                </div>
            </section>

            {/* 4. Posts/Social */}
            <section>
                <div className="flex justify-between items-center mb-6">
                    <h1 className="text-xs font-bold uppercase tracking-widest text-purple-600">Community Posts</h1>
                    <Link to="/post" className="bg-purple-600 text-white px-4 py-2 rounded-full text-sm font-bold hover:bg-purple-700 transition">
                        + Create Post
                    </Link>
                </div>
                
                <div className="space-y-6">
                    {data.Posts.sort((a,b) => Date.parse(b.createdAt) - Date.parse(a.createdAt)).map((p, i) => (
                        <div key={i} className="flex gap-4 items-start">
                            <div className="w-10 h-10 rounded-full bg-purple-100 flex-shrink-0 flex items-center justify-center text-purple-600 font-bold">
                                {p.name[0]}
                            </div>
                            <div className="flex flex-col">
                                <span className="font-bold text-sm text-gray-900">{p.name}</span>
                                <p className="text-gray-600 text-sm mt-1">{p.body}</p>
                            </div>
                        </div>
                    ))}
                </div>
            </section>
        </div>
    );
}