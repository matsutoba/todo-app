"use client";

import PageHeader from "../components/PageHeader";

export const MainLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="flex h-screen">
      <main className="bg-slate-50 flex-1 overflow-auto">
        <PageHeader />
        <div className="p-4">{children}</div>
      </main>
    </div>
  );
};

export default MainLayout;
