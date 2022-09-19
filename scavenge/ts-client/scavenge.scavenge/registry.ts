import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCommitSolution } from "./types/scavenge/tx";
import { MsgRevealSolution } from "./types/scavenge/tx";
import { MsgSubmitScavenge } from "./types/scavenge/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/scavenge.scavenge.MsgCommitSolution", MsgCommitSolution],
    ["/scavenge.scavenge.MsgRevealSolution", MsgRevealSolution],
    ["/scavenge.scavenge.MsgSubmitScavenge", MsgSubmitScavenge],
    
];

export { msgTypes }