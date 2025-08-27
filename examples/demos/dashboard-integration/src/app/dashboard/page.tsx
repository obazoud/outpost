"use client";

import { useSession } from "next-auth/react";
import { useSearchParams } from "next/navigation";
import { useEffect, useState } from "react";
import { useOverview } from "@/hooks/useOverview";
import OverviewStats from "@/components/dashboard/OverviewStats";
import RecentActivity from "@/components/dashboard/RecentActivity";

export default function DashboardPage() {
  const { data: session } = useSession();
  const { data: overview, loading, error } = useOverview();
  const searchParams = useSearchParams();
  const [showWelcome, setShowWelcome] = useState(false);

  useEffect(() => {
    if (searchParams.get("welcome") === "true") {
      setShowWelcome(true);
      // Auto-hide the welcome message after 5 seconds
      setTimeout(() => setShowWelcome(false), 5000);
    }
  }, [searchParams]);

  if (loading) {
    return (
      <div className="space-y-6">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Dashboard</h1>
          <p className="text-gray-600">
            Loading your event destinations overview...
          </p>
        </div>
        <div className="animate-pulse">
          <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
            {[...Array(4)].map((_, i) => (
              <div
                key={i}
                className="bg-white p-6 rounded-lg border border-gray-200"
              >
                <div className="h-4 bg-gray-200 rounded w-3/4 mb-2"></div>
                <div className="h-8 bg-gray-200 rounded w-1/2"></div>
              </div>
            ))}
          </div>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="space-y-6">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Dashboard</h1>
          <p className="text-red-600">Error loading dashboard: {error}</p>
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-gray-900">Dashboard</h1>
        <p className="text-gray-600">
          Welcome back, {session?.user?.name || session?.user?.email}
        </p>
      </div>

      {showWelcome && (
        <div className="bg-green-50 border border-green-200 rounded-md p-4">
          <div className="flex">
            <div className="flex-shrink-0">
              <svg
                className="h-5 w-5 text-green-400"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fillRule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                  clipRule="evenodd"
                />
              </svg>
            </div>
            <div className="ml-3">
              <h3 className="text-sm font-medium text-green-800">
                Welcome to your dashboard! 🎉
              </h3>
              <div className="mt-2 text-sm text-green-700">
                <p>
                  Your account has been created and you're ready to start
                  managing event destinations. Click "Event Destinations" to get
                  started!
                </p>
              </div>
            </div>
          </div>
        </div>
      )}

      {overview && (
        <>
          <OverviewStats stats={overview.stats} />

          <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <RecentActivity events={overview.recentEvents} />

            <div className="bg-white p-6 rounded-lg border border-gray-200">
              <h3 className="text-lg font-medium text-gray-900 mb-4">
                Destinations ({overview.destinations.length})
              </h3>
              {overview.destinations.length === 0 ? (
                <p className="text-sm text-gray-500">
                  No destinations configured yet.{" "}
                  <a
                    href="/dashboard/event-destinations"
                    className="text-blue-600 hover:text-blue-500"
                  >
                    Add your first destination →
                  </a>
                </p>
              ) : (
                <div className="space-y-2">
                  {overview.destinations.map((dest) => (
                    <a
                      key={dest.id}
                      href={`/dashboard/event-destinations/destinations/${dest.id}`}
                      className="flex items-center justify-between p-2 bg-gray-50 rounded hover:bg-gray-100 transition-colors cursor-pointer"
                    >
                      <div>
                        <span className="text-sm font-medium text-gray-900">
                          {dest.type}
                        </span>
                        {dest.config.url && (
                          <p className="text-xs text-gray-600">
                            {dest.config.url}
                          </p>
                        )}
                      </div>
                      <span
                        className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                          dest.enabled
                            ? "bg-green-100 text-green-800"
                            : "bg-gray-100 text-gray-800"
                        }`}
                      >
                        {dest.enabled ? "Active" : "Inactive"}
                      </span>
                    </a>
                  ))}
                  <div className="pt-2">
                    <a
                      href="/dashboard/event-destinations"
                      className="text-sm text-blue-600 hover:text-blue-500"
                    >
                      Manage all destinations →
                    </a>
                  </div>
                </div>
              )}
            </div>
          </div>
        </>
      )}
    </div>
  );
}
