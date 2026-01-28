import { getHome } from '../../hooks/home';
import { getGuest } from '../../hooks/identity';
import { useParams, Link, useNavigate } from 'react-router-dom';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useState } from 'react';
import { Header } from '../Header';

export const Home = () => {
    const navigate = useNavigate()
    const { code } = useParams();
    const queryClient = useQueryClient();
    if (code === undefined) {
        navigate('/')
    }
    // 1. State for inline editing
    const [editingPostId, setEditingPostId] = useState<string | null>(null);
    const [editBody, setEditBody] = useState("");

    if (code === undefined) {
        navigate('/');
    }

    const { data, isLoading, error } = getHome(code);
    const guest_id = getGuest(code);

    if (guest_id === null) {
        navigate(`/login/${code}`);
    }

    // 2. Mutations for Update and Delete
    const deletePostMutation = useMutation({
        mutationFn: async (postId: number) => {
            const res = await fetch(`${import.meta.env.VITE_API_URL}/post/${code}`, {
                method: 'DELETE',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ id: postId }),
            });
            if (!res.ok) throw new Error('Could not delete post');
        },
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['home', code] });
        }
    });

    const updatePostMutation = useMutation({
        mutationFn: async (postId: number) => {
            const res = await fetch(`${import.meta.env.VITE_API_URL}/post/${code}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ id: postId, body: editBody }),
            });
            if (!res.ok) throw new Error('Could not update post');
        },
        onSuccess: () => {
            setEditingPostId(null);
            queryClient.invalidateQueries({ queryKey: ['home', code] });
        }
    });
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
            <Header />
            {/* 1. About Section */}
            <section>
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
                    {data.Announcements.sort((a,b) => Date.parse(b.created_at) - Date.parse(a.created_at)).map((a, i) => (
                        <div key={i} className="p-4 bg-orange-50 rounded-xl border border-orange-100">
                            <h3 className="font-bold text-gray-900">{a.header}</h3>
                            <div className="text-gray-700 mt-1">{a.body}</div>
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
                                <Link to={`/guest/${code}`} className="text-xs bg-white border px-3 py-1 rounded hover:bg-gray-100 transition shadow-sm">
                                    Edit Profile
                                </Link>
                            )}
                        </div>
                    ))}
                </div>
            </section>

            {/* 4. Community Posts with Edit/Delete */}
            <section>
                <div className="flex justify-between items-center mb-6">
                    <h1 className="text-xs font-bold uppercase tracking-widest text-purple-600">Community Posts</h1>
                    <Link to={`/post/${code}`} className="bg-purple-600 text-white px-4 py-2 rounded-full text-sm font-bold hover:bg-purple-700 transition">
                        + Create Post
                    </Link>
                </div>
                
                <div className="space-y-6">
                    {data.Posts.sort((a,b) => Date.parse(b.created_at) - Date.parse(a.created_at)).map((p, i) => (
                        <div key={i} className="flex gap-4 items-start group">
                            <div className="w-10 h-10 rounded-full bg-purple-100 flex-shrink-0 flex items-center justify-center text-purple-600 font-bold">
                                {p.name[0]}
                            </div>
                            <div className="flex-1 flex flex-col">
                                <div className="flex justify-between items-center">
                                    <span className="font-bold text-sm text-gray-900">{p.name}</span>
                                    
                                    {/* 3. Show controls only if the guest_id matches */}
                                    {p.guest_id === guest_id && editingPostId !== p.id && (
                                        <div className="flex gap-3 opacity-0 group-hover:opacity-100 transition-opacity">
                                            <button 
                                                onClick={() => { setEditingPostId(p.id); setEditBody(p.body); }}
                                                className="text-xs text-blue-600 font-semibold hover:underline"
                                            >
                                                Edit
                                            </button>
                                            <button 
                                                onClick={() => { if(confirm("Delete post?")) deletePostMutation.mutate(p.id); }}
                                                className="text-xs text-red-600 font-semibold hover:underline"
                                            >
                                                Delete
                                            </button>
                                        </div>
                                    )}
                                </div>

                                {editingPostId === p.id ? (
                                    <div className="mt-2 space-y-2">
                                        <textarea 
                                            className="w-full p-2 text-sm border rounded-lg focus:ring-2 focus:ring-purple-500 outline-none"
                                            value={editBody}
                                            onChange={(e) => setEditBody(e.target.value)}
                                        />
                                        <div className="flex gap-2">
                                            <button 
                                                onClick={() => updatePostMutation.mutate(p.id)}
                                                className="bg-purple-600 text-white px-3 py-1 rounded text-xs font-bold"
                                            >
                                                Save
                                            </button>
                                            <button 
                                                onClick={() => setEditingPostId(null)}
                                                className="bg-gray-100 text-gray-600 px-3 py-1 rounded text-xs"
                                            >
                                                Cancel
                                            </button>
                                        </div>
                                    </div>
                                ) : (
                                    <p className="text-gray-600 text-sm mt-1">{p.body}</p>
                                )}
                            </div>
                        </div>
                    ))}
                </div>
            </section>
        </div>
    );
}