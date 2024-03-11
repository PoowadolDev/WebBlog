export default function Page() {
    return (
        <div className="container mx-auto w-1/3 mt-40 bg-white text-black font-sans rounded-3xl">
            <div className="flex justify-center box-border">
                <div>
                    <h2 className="text-5xl font-bold py-6 px-12 ">Sign Up</h2>
                    <form action="#">
                    <div className="mt-5">
                        <input type="text" placeholder="Username" className="border"/>
                    </div>
                    <div className="mt-5">
                        <input type="email" placeholder="Email" className="border"/>
                    </div>
                    <div className="mt-5">
                        <i className="fa fa-user icon"></i>
                        <input type="password" placeholder="password" className="border"/>
                    </div>
                    <div className="mt-5">
                        <input type="checkbox" placeholder="box" className="mr-2"/>
                        <span>I agree with the terms</span>
                    </div>
                    <div className="mt-5 mb-5">
                        <button type="submit" className="bg-indigo-600 text-white px-7 py-2 rounded-3xl">Sign Up</button>
                    </div>
                </form>
                </div>


            </div>
        </div>
    )
}