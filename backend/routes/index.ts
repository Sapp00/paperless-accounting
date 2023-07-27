import { Prisma, PrismaClient } from "@prisma/client";

const prisma = new PrismaClient();

async function main() {
    const entry = await prisma.expense.create({
        data: {
            price: 12.89,
            paid: false
        }
    });
    console.log(entry);
}

main()
    .then(async() => {
        await prisma.$disconnect();
    })
    .catch(async(e) => {
        console.error(e);
        await prisma.$disconnect();
        process.exit();
    });

