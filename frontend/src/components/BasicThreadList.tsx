import ThreadCard from "./ThreadCard";
import React, { useEffect, useState } from "react";

interface Thread {
    id: number;
    title: string;
    authorId: number;
    CategoryID: number;
    content: string;
    createdAt: string;
}

const BasicThreadList: React.FC = () => {
    const [threads, setThreads] = useState<Thread[]>([]);
    const [currentUserId, setCurrentUserId] = useState<string | null>(null);

    useEffect(() => {
        
        const userId = localStorage.getItem("userId");
        setCurrentUserId(userId);

        fetch("http://localhost:8080/api/v1/threads")
            .then((response) => response.json())
            .then((data) => {
                const sortedThreads = data.sort(
                    (a: Thread, b: Thread) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime(),
                );
                setThreads(sortedThreads);
            })
            .catch((error) => console.error("Error fetching threads:", error));
    }, []);

    return (
        <div style={{ width: "70vw", margin: "auto", paddingTop: 20 }}>
            {threads.length > 0 ? (
                threads.map((thread) => (
                    <ThreadCard
                        key={thread.id}
                        categoryId={thread.CategoryID}
                        title={thread.title}
                        authorId={thread.authorId}
                        content={thread.content}
                        threadId={thread.id}
                        createdAt={thread.createdAt}
                        buttonNeeded={true}
                        isCreator={currentUserId === thread.authorId.toString()}
                    />
                ))
            ) : (
                <p>{"No threads have been posted. Be the first to share your thoughts!"}</p>
            )}
        </div>
    );
};

export default BasicThreadList;
