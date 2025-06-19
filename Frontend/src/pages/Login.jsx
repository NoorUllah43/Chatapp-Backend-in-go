import { useState } from "react";
import { IoChatbubbleEllipsesOutline, IoChatbubbleEllipsesSharp } from "react-icons/io5";
import { MdEmail, MdPassword } from "react-icons/md";
import { FaUser } from "react-icons/fa";
import { IoEyeOutline, IoEyeOffOutline } from "react-icons/io5";
import { motion } from "motion/react";



const Login = () => {
    const [formType, setFormType] = useState('login');
    const [name, setName] = useState('')
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('')
    const [isSubmitting, setIsSubmitting] = useState(false)
    const [passwordVisibility, setPasswordVisibility] = useState(false)

    const submitForm = async (e) => {
        e.preventDefault()
        setIsSubmitting(true)
        if (formType == "login") {
            console.log(email, password)
        }
        else {
            console.log(name, email, password)
        }
        setIsSubmitting(false)

    }



    return (
        <div className='flex items-center justify-between '>
            <div className='flex flex-col items-center justify-center gap-2 w-1/2 h-screen bg-gradient-to-b from-purple-500 to-blue-500 text-white '>
                <motion.div animate={{
                    y: [0, -40, 0],
                    rotate: [0, -10, 0],
                }}
                    transition={{
                        duration: 5,
                        repeat: Infinity,
                        repeatType: "loop",
                        ease: "easeInOut",
                    }} className='flex items-center justify-center text-9xl font-bold '>
                    <IoChatbubbleEllipsesOutline />
                </motion.div>
                <h2 className="text-3xl font-bold">Wlecome to ChatVerse</h2>
                <p className="text-sm text-gray-200">connect and chat with anyone anywhere</p>
            </div>
            <div className="flex items-center justify-center w-1/2 h-screen">
                <div className="flex flex-col items-center justify-center gap-2 h-full w-[400px]">
                    <h2 className="text-3xl font-bold flex gap-2"><span className="text-purple-500 text-4xl"><IoChatbubbleEllipsesSharp /></span> ChatVerse</h2>
                    <p className="text-gray-600">{formType == 'login' ? 'Sign to continue' : 'Create an account'}</p>
                    <form onSubmit={submitForm} className="flex flex-col items-center justify-center w-full gap-4 my-5">

                        {/* Name input  */}
                        {formType == "signup" && (
                            <div className="flex gap-2 items-center border-2 border-gray-300 rounded-md p-2 h-[42px] w-[400px] focus-within:ring-2 focus-within:ring-purple-500 ">
                                <label htmlFor="name" className="text-gray-400 text-2xl w-[30px]"><FaUser /></label>
                                <input type="name" name="name" placeholder="Enter your name" id="name" required value={name} onChange={(e) => setName(e.target.value)} className="outline-none h-full w-[calc(100%-30px)] " />
                            </div>
                        )}

                        {/* email input  */}
                        <div className="flex gap-2 items-center border-2 border-gray-300 rounded-md p-2 h-[42px] w-[400px] focus-within:ring-2 focus-within:ring-purple-500">
                            <label htmlFor="email" className="text-gray-400 text-2xl w-[30px]"><MdEmail /></label>
                            <input type="email" name="email" placeholder="Enter your email" id="email" required value={email} onChange={(e) => setEmail(e.target.value)} className="outline-none h-full w-[calc(100%-30px)] " />
                        </div>
                        {/* password input  */}
                        <div className="flex gap-2 items-center border-2 border-gray-300 rounded-md p-2 h-[42px] w-[400px] focus-within:ring-2 focus-within:ring-purple-500">
                            <label htmlFor="password" className="text-gray-400 text-2xl w-[30px]"><MdPassword /> </label>
                            <input type={passwordVisibility ? "text" : "password"} name="password" placeholder="Enter your password" id="password" required value={password} onChange={(e) => setPassword(e.target.value)} className="outline-none h-full w-[calc(100%-55px)] " />
                            <button type="button" onClick={() => setPasswordVisibility(!passwordVisibility)} className="w-5 h-full flex items-center justify-center text-gray-500" >{passwordVisibility ? <IoEyeOutline /> : <IoEyeOffOutline />}</button>
                        </div>
                        {formType == "login" && (
                            <p className="text-purple-500 text-right w-full ">Forget password?</p>
                        )}
                        <button type="submit" disabled={isSubmitting} className={`bg-purple-500 hover:bg-purple-600 duration-200 text-white font-bold py-2 px-4 rounded-md h-[42px] w-full flex items-center justify-center ${isSubmitting ? 'opacity-50' : ''}`}>
                            {
                                isSubmitting ? (
                                    <p className="w-[20px] h-[20px] rounded-full border-2 border-white border-r-gray-400 animate-spin"></p>
                                ) : (
                                    formType === 'login' ? 'Login' : 'Sign Up'
                                )
                            }

                        </button>
                    </form>
                    {formType == "login" ? (
                        <p className="text-gray-500 my-3 ">Don't have an account? <span onClick={() => setFormType('signup')} className="text-purple-500 cursor-pointer"> Sign Up</span></p>

                    ) : (
                        <p className="text-gray-500 my-3 ">Already have an account? <span onClick={() => setFormType('login')} className="text-purple-500 cursor-pointer"> Login</span></p>

                    )}
                </div>
            </div>
        </div>
    )
}

export default Login