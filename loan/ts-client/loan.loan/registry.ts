import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRequestLoan } from "./types/loan/tx";
import { MsgApproveLoan } from "./types/loan/tx";
import { MsgRepayLoan } from "./types/loan/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/loan.loan.MsgRequestLoan", MsgRequestLoan],
    ["/loan.loan.MsgApproveLoan", MsgApproveLoan],
    ["/loan.loan.MsgRepayLoan", MsgRepayLoan],
    
];

export { msgTypes }