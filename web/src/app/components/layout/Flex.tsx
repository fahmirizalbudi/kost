const Flex = ({ className, children }: { className?: string, children: React.ReactNode }) => {
  return <div style={{ display: "flex" }} className={className}>{children}</div>
}

export default Flex
