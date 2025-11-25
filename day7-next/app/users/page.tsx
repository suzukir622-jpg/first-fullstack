"use client";

import { useEffect, useState } from "react";

type User = {
  name: string;
  role: string;
};

export default function UsersPage() {
  const [users, setUsers] = useState<User[]>([]);
  const [name, setName] = useState("");
  const [role, setRole] = useState("");

  // ★ loadUsers
  const loadUsers = () => {
    fetch("http://127.0.0.1:8080/users")
      .then((res) => res.json())
      .then((data) => setUsers(data))
      .catch((err) => console.error("fetch error:", err));
  };

  // ★ addUser
  const addUser = () => {
    fetch("http://127.0.0.1:8080/add", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name, role }),
    })
      .then((res) => res.json())
      .then((data) => setUsers(data))
      .catch((err) => console.error("add error:", err));
  };

  // ★ deleteUser
  const deleteUser = (index: number) => {
    fetch("http://127.0.0.1:8080/delete", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ index }),
    })
      .then((res) => res.json())
      .then((data) => setUsers(data))
      .catch((err) => console.error("delete error:", err));
  };

  // ★ 初回読み込み
  useEffect(() => {
    loadUsers();
  }, []);

  return (
    <div>
      <h1>Users</h1>

      <div>
        <input
          placeholder="name"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <input
          placeholder="role"
          value={role}
          onChange={(e) => setRole(e.target.value)}
        />
        <button onClick={addUser}>Add</button>
      </div>

      <div>
        {users.map((user, i) => (
          <div key={i}>
            {user.name} / {user.role}
            <button onClick={() => deleteUser(i)}>Delete</button>
          </div>
        ))}
      </div>
    </div>
  );
}


