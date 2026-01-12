import { useParams, Link } from 'react-router-dom';
import { postUnsubscribeAll } from '../../hooks/unsubscribe';

export const UnsubscribeAll = () => {
    const { email } = useParams();

    // Reusable Error Component for consistency
    const ErrorView = ({ message }: { message: string }) => (
        <div className="flex items-center justify-center min-h-screen bg-gray-50 p-4">
            <div className="max-w-md w-full bg-white p-8 rounded-2xl shadow-sm border border-red-100 text-center">
                <div className="w-16 h-16 bg-red-50 text-red-500 rounded-full flex items-center justify-center mx-auto mb-4 text-2xl font-bold">
                    !
                </div>
                <h1 className="text-xl font-bold text-gray-900 mb-2">Action Failed</h1>
                <p className="text-gray-600 mb-6">{message}</p>
                <a href="mailto:support@partiburo.com" className="text-blue-600 font-medium hover:underline">
                    Contact Support
                </a>
            </div>
        </div>
    );

    if (email === undefined) {
        return <ErrorView message="No email address found. Please report this to support@partiburo.com" />;
    }

    const { data, isLoading, error } = postUnsubscribeAll(email);

    if (isLoading) {
        return (
            <div className="flex items-center justify-center min-h-screen bg-gray-50">
                <div className="flex flex-col items-center gap-4">
                    <div className="w-12 h-12 border-4 border-blue-600 border-t-transparent rounded-full animate-spin"></div>
                    <span className="text-gray-600 font-medium tracking-wide">Processing your request...</span>
                </div>
            </div>
        );
    }

    if (error || data?.Code !== 200) {
        return <ErrorView message="Failed to unsubscribe from all communications. Please contact support@partiburo.com" />;
    }

    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-50 p-4">
            <div className="max-w-md w-full bg-white p-10 rounded-2xl shadow-xl border border-gray-100 text-center animate-in fade-in zoom-in duration-300">
                {/* Visual indicator for 'All/Global' action */}
                <div className="w-20 h-20 bg-gray-200 text-gray-600 rounded-full flex items-center justify-center mx-auto mb-6 text-4xl">
                    ✕
                </div>
                
                <h1 className="text-2xl font-extrabold text-gray-900 mb-4 tracking-tight">Globally Unsubscribed</h1>
                
                <p className="text-gray-600 mb-4 leading-relaxed">
                    You have been successfully unsubscribed from <strong>all future communication</strong> from Partiburo.
                </p>
                
                <div className="bg-blue-50 p-4 rounded-xl mb-8 border border-blue-100">
                    <p className="text-sm text-blue-800">
                        To undo this or if this was a mistake, please reach out to 
                        <a href="mailto:support@partiburo.com" className="font-bold underline ml-1">support@partiburo.com</a>.
                    </p>
                </div>

                <Link 
                    to="/" 
                    className="inline-block w-full bg-gray-900 text-white font-bold py-3 px-6 rounded-xl hover:bg-black transition-all shadow-lg shadow-gray-200 active:scale-[0.98]"
                >
                    Back to Home
                </Link>
            </div>
        </div>
    );
};