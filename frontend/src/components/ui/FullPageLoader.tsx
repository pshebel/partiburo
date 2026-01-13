export const FullPageLoader = () => {
  return (
    <div className="min-h-screen w-full flex flex-col items-center justify-center bg-gray-50">
      <div className="relative flex items-center justify-center">
        {/* Outer Spinning Ring */}
        <div className="w-20 h-20 border-4 border-blue-100 border-t-blue-600 rounded-full animate-spin"></div>
        
        {/* Inner Pulsing Dot or Icon */}
        <div className="absolute w-10 h-10 bg-blue-600 rounded-full animate-pulse opacity-20"></div>
      </div>
      
      <h2 className="mt-6 text-sm font-bold uppercase tracking-widest text-gray-400 animate-pulse">
        Partiburo
      </h2>
    </div>
  );
};