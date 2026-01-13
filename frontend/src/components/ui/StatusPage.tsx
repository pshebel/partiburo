// src/components/ui/StatusPage.tsx
import { Link } from 'react-router-dom';

interface StatusPageProps {
  status: 'loading' | 'error' | 'success';
  title: string;
  message: string;
  icon?: string;
  buttonText?: string;
  errorContact?: string;
}

export const StatusPage = ({ status, title, message, icon, buttonText = "Back home", errorContact }: StatusPageProps) => {
  if (status === 'loading') {
    return (
      <div className="flex items-center justify-center min-h-screen bg-gray-50">
        <div className="flex flex-col items-center gap-4">
          <div className="w-12 h-12 border-4 border-blue-600 border-t-transparent rounded-full animate-spin"></div>
          <span className="text-gray-600 font-medium">Processing...</span>
        </div>
      </div>
    );
  }

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-50 p-4">
      <div className="max-w-md w-full bg-white p-10 rounded-2xl shadow-xl border border-gray-100 text-center animate-in fade-in zoom-in duration-300">
        <div className={`w-20 h-20 rounded-full flex items-center justify-center mx-auto mb-6 text-4xl ${status === 'error' ? 'bg-red-50 text-red-500' : 'bg-gray-100 text-gray-500'}`}>
          {icon || (status === 'error' ? '!' : '✓')}
        </div>
        <h1 className="text-2xl font-extrabold text-gray-900 mb-2">{title}</h1>
        <p className="text-gray-600 mb-8 leading-relaxed">{message}</p>
        
        {status === 'error' && errorContact && (
           <div className="mb-6 text-sm text-gray-500">Please report this to <a href={`mailto:${errorContact}`} className="text-blue-600 underline">{errorContact}</a></div>
        )}

        <Link to="/" className="inline-block w-full bg-gray-900 text-white font-bold py-3 px-6 rounded-xl hover:bg-black transition-all">
          {buttonText}
        </Link>
      </div>
    </div>
  );
};