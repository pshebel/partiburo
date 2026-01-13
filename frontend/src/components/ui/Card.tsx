// src/components/ui/Card.tsx
export const Card = ({ title, subtitle, children, accentColor = "blue" }: { title: string, subtitle: string, children: React.ReactNode, accentColor?: string }) => {
  const accentClasses: Record<string, string> = {
    blue: "text-blue-600",
    purple: "text-purple-600",
    orange: "text-orange-600"
  };

  return (
    <section className="bg-white p-8 rounded-2xl shadow-sm border border-gray-100">
      <div className="border-b pb-4 mb-6">
        <h1 className={`text-xs font-bold uppercase tracking-widest ${accentClasses[accentColor]} mb-1`}>{title}</h1>
        <h2 className="text-2xl font-extrabold text-gray-900 leading-tight">{subtitle}</h2>
      </div>
      {children}
    </section>
  );
};