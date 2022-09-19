import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRevealSolution } from "./types/scavenge/tx";
import { MsgSubmitScavenge } from "./types/scavenge/tx";
import { MsgCommitSolution } from "./types/scavenge/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/scavenge.scavenge.MsgRevealSolution", MsgRevealSolution],
    ["/scavenge.scavenge.MsgSubmitScavenge", MsgSubmitScavenge],
    ["/scavenge.scavenge.MsgCommitSolution", MsgCommitSolution],
    
];

export { msgTypes }