\hline 
	SUCCESS             & 0 &  & 成功 \\ 
\hline
	ERROR\_BAD\_REQUEST   & 1002 & & 参数错误 \\ 
\hline
	ERROR\_TOKEN\_INVALID & 1003 & & 无效的Token \\ 
\hline
	ERROR\_USER\_NOT\_FOUND & 2001 & & 该用户不存在 \\ 
\hline
	ERROR\_USER\_DUPLICATE & 2002 & & 存在重复用户 \\ 
\hline
	ERROR\_USER\_WRONG\_PWD & 2003 & & 密码错误 \\ 
\hline
	ERROR\_USER\_UNAUTHEN  & 2004 & & 用户未登录 \\ 
\hline
	ERROR\_USER\_OFFLINE   & 2005 & & 用户未在线 \\ 
\hline
	ERROR\_GROUP\_DUPLICATE & 2051 & & 存在重复群组 \\ 
\hline
	ERROR\_GROUP\_NOTFOUND  & 2052 & & 该群组不存在 \\ 
\hline
	ERROR\_SERVER\_COMMON           &  3000 & 一般性错误 & 服务器内部错误 \\ 
\hline
	ERROR\_MQ                      &  3001 & 消息队列错误 & 服务器内部错误 \\ 
\hline
	ERROR\_REDIS                   &  3003 & Redis错误 & 服务器内部错误 \\ 
\hline
	ERROR\_SEQID                   &  3004 & 序列号错误 & 服务器内部错误 \\ 
\hline
	ERROR\_GW\_USER\_NOT\_FOUND  &  3005 & 用户不在此网关上 & 服务器内部错误 \\ 
\hline
	ERROR\_IO                      &  3006 & socket io 错误 & 服务器内部错误 \\ 
\hline
	ERROR\_MARSHALL                &  3007 & 序列化错误 & 服务器内部错误 \\ 
\hline
	ERROR\_PART                    &  3008 & 部分错误 & 服务器内部错误 \\ 
\hline
	ERROR\_DB                      &  3010 & 数据库错误 & 数据库繁忙 \\ 
\hline
	ERROR\_CASSANDRA               &  3011 & Cassandra错误 & 数据库繁忙 \\ 
\hline
	ERROR\_DGRAPH               &  3012 & Dgraph错误 & 数据库繁忙 \\ 
\hline
	ERROR\_DISCOVER               &  3013 & 服务发现错误 & 服务器内部错误 \\ 
\hline
	ERROR\_BC\_GAS         & 4001 & gas 相关错误 & 区块链错误 \\ 
\hline
	ERROR\_BC\_NONCE       & 4002 & nonce 相关错误 & 区块链错误 \\ 
\hline
	ERROR\_BC\_TRANSACTION & 4003 & 交易错误 & 区块链错误 \\ 
\hline
	ERROR\_UNKNOWN              & 9001 &  & 未知错误 \\ 
\hline
	ERROR\_UNIMPLEMENTED & 9002 &  & 功能未实现 \\