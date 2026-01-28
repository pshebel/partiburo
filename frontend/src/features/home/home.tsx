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
        if (error || !data) {
            return (
                <div className="min-h-screen bg-gray-50 flex items-center justify-center p-6">
                    <div className="max-w-md w-full p-6 text-center bg-white rounded-2xl shadow-sm border border-red-100">
                    <div className="text-red-500 text-4xl mb-4">⚠️</div>
                    <h2 className="text-xl font-bold text-gray-900 mb-2">Party Not Found</h2>
                    <p className="text-gray-600 mb-6">
                        {error?.message || "We couldn't find a party with that code. Please check your link and try again."}
                    </p>
                    <button 
                        onClick={() => navigate('/')}
                        className="w-full py-2 bg-gray-900 text-white rounded-lg font-medium hover:bg-gray-800 transition-colors"
                    >
                        Go Back Home
                    </button>
                    </div>
                </div>
            );
        }
    }

    return (
        <div className="max-w-4xl mx-auto p-6 space-y-12 bg-white min-h-screen">
            <Header />
            {/* 1. About Section */}
            {/* 1. Hero Section - Centered for Balance */}
            <section className="text-center py-8 border-b border-gray-200">
                <span className="text-xs font-bold uppercase tracking-[0.2em] text-blue-600">
                    You're Invited
                </span>
                <h1 className="text-5xl font-black text-gray-900 mt-2 mb-6 tracking-tight">
                    {data.Title}
                </h1>
                <p className="max-w-2xl mx-auto text-xl text-gray-600 leading-relaxed mb-8 whitespace-pre-line">
                    {data.Description}
                </p>
                
                {/* Centered Pill Detail Bar */}
                <div className="inline-flex flex-wrap justify-center gap-6 px-8 py-4 bg-white rounded-2xl shadow-sm border border-gray-100 text-sm font-semibold text-gray-700">
                    <div className="flex items-center gap-2">📅 {data.Date}</div>
                    <div className="hidden sm:block border-r border-gray-200 h-4" />
                    <div className="flex items-center gap-2">⏰ {data.Time}</div>
                    <div className="hidden sm:block border-r border-gray-200 h-4" />
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