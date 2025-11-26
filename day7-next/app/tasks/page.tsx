"use client";

import { useEffect, useState } from "react";

type Task = {
  id: number;
  name: string;
};

export default function TasksPage() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    async function loadTasks() {
      try {
        const res = await fetch("http://localhost:8080/tasks");
        if (!res.ok) {
          throw new Error(`failed: ${res.status}`);
        }
        const data: Task[] = await res.json();
        setTasks(data);
      } catch (e: any) {
        setError(e.message ?? "error");
      } finally {
        setLoading(false);
      }
    }

    loadTasks();
  }, []);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <div>
      <h1>Tasks</h1>
      <ul>
        {tasks.map((t) => (
          <li key={t.id}>
            {t.id}: {t.name}
          </li>
        ))}
      </ul>
    </div>
  );
}
