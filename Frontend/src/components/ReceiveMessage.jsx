
const ReceiveMessage = ({message, time}) => {

  return (
     <div className='w-auto h-auto flex flex-col items-start p-3'>
      <div className='bg-gray-200 w-auto h-auto inline-block p-2 rounded-t-xl rounded-e-xl '>
        <p className='w-auto h-auto inline-block'>{message}</p>
      </div>
      <div className='w-auto h-auto inline-block px-2'>
        <p className='w-auto h-auto inline-block text-[10px] md:text-[12px] '>{time}</p>
      </div>
    </div>
  )
}

export default ReceiveMessage
