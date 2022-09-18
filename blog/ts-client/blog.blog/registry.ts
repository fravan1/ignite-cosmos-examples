import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgDeleteComment } from "./types/blog/tx";
import { MsgCreateComment } from "./types/blog/tx";
import { MsgCreatePost } from "./types/blog/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/blog.blog.MsgDeleteComment", MsgDeleteComment],
    ["/blog.blog.MsgCreateComment", MsgCreateComment],
    ["/blog.blog.MsgCreatePost", MsgCreatePost],
    
];

export { msgTypes }