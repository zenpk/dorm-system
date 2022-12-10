// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type {NextApiRequest, NextApiResponse} from 'next'

type Data = {
    msg: string
};

let alertMessage = "";

export default function handler(
    req: NextApiRequest,
    res: NextApiResponse<Data>
) {
    switch (req.method) {
        case "GET":
            res.json({msg: alertMessage});
            alertMessage = "";
            return;
        case "POST":
            alertMessage = req.body.toString();
            res.json({msg: "ok"});
            return;
        default:
            res.json({msg: "no correspondent method"});
            return;
    }
}
