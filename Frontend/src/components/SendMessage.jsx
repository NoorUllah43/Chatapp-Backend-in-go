
const SendMessage = ({message,time }) => {

  return (
    <div className='w-auto h-auto flex flex-col items-end p-3'>
      <div className='bg-[#4A7BF7] text-white w-auto h-auto inline-block p-2 rounded-t-xl rounded-s-xl '>
        <p className='w-auto h-auto inline-block'>{message}</p>
      </div>
      <div className='w-auto h-auto inline-block px-2'>
        <p className='w-auto h-auto inline-block text-[10px] md:text-[12px]'>{time}</p>
      </div>
    </div>

  )
}

export default SendMessage
