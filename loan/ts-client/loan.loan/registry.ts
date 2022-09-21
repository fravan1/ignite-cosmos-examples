import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgApproveLoan } from "./types/loan/tx";
import { MsgRequestLoan } from "./types/loan/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/loan.loan.MsgApproveLoan", MsgApproveLoan],
    ["/loan.loan.MsgRequestLoan", MsgRequestLoan],
    
];

export { msgTypes }