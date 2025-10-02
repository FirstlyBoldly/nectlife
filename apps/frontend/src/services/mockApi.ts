export type PostSummary = {
    id: number;
    title: string;
};

export type Post = {
    content: string;
} & PostSummary;

const range = (n: number): number[] => [...Array(n).keys()];
const MOCK_POSTS: Post[] = range(20).map((i) => {
    return {
        id: i,
        title: `Testy Mc Testface ${i + 1}`,
        content: ".",
    };
});

export const api = {
    fetchPosts: (): Promise<PostSummary[]> => {
        console.log("API: Fetching all post summaries...");
        return new Promise((resolve) => {
            setTimeout(() => {
                const summaries = MOCK_POSTS;
                console.log("API: Responded with post summaries.");
                resolve(summaries);
            }, 500);
        });
    },

    fetchPostsById: (id: number): Promise<Post | undefined> => {
        console.log(`API: Fetching post with id: ${id}...`);
        return new Promise<Post | undefined>((resolve) => {
            setTimeout(() => {
                const post = MOCK_POSTS.find((p) => p.id === id);
                console.log(`API: Responded for post with id: ${id}.`);
                resolve(post);
            }, 500);
        });
    },
};

export const MockImageUpload = async (file: File): Promise<string> => {
    await new Promise((resolve) => setTimeout(resolve, 1000));
    const imageURL = `https://picsum.photos/seed/${file.name}/800/400`;
    return imageURL;
};
