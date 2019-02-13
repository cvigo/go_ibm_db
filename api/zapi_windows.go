// mksyscall_windows.pl api.go
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package api

import "unsafe"
import "syscall"
import "os"

var (
	mododbc32 = syscall.NewLazyDLL(GetDllName())

	procSQLAllocHandle     = mododbc32.NewProc("SQLAllocHandle")
	procSQLBindCol         = mododbc32.NewProc("SQLBindCol")
	procSQLBindParameter   = mododbc32.NewProc("SQLBindParameter")
	procSQLCloseCursor     = mododbc32.NewProc("SQLCloseCursor")
	procSQLDescribeColW    = mododbc32.NewProc("SQLDescribeColW")
	procSQLDescribeParam   = mododbc32.NewProc("SQLDescribeParam")
	procSQLDisconnect      = mododbc32.NewProc("SQLDisconnect")
	procSQLDriverConnectW  = mododbc32.NewProc("SQLDriverConnectW")
	procSQLEndTran         = mododbc32.NewProc("SQLEndTran")
	procSQLExecute         = mododbc32.NewProc("SQLExecute")
	procSQLFetch           = mododbc32.NewProc("SQLFetch")
	procSQLFreeHandle      = mododbc32.NewProc("SQLFreeHandle")
	procSQLGetData         = mododbc32.NewProc("SQLGetData")
	procSQLGetDiagRecW     = mododbc32.NewProc("SQLGetDiagRecW")
	procSQLNumParams       = mododbc32.NewProc("SQLNumParams")
	procSQLNumResultCols   = mododbc32.NewProc("SQLNumResultCols")
	procSQLPrepareW        = mododbc32.NewProc("SQLPrepareW")
	procSQLRowCount        = mododbc32.NewProc("SQLRowCount")
	procSQLSetEnvAttr      = mododbc32.NewProc("SQLSetEnvAttr")
	procSQLSetConnectAttrW = mododbc32.NewProc("SQLSetConnectAttrW")
	procSQLColAttribute    = mododbc32.NewProc("SQLColAttribute")
)

func GetDllName() string {
	if winArch := os.Getenv("PROCESSOR_ARCHITECTURE"); winArch == "x86" {
		return "db2cli.dll"
	} else {
		return "db2cli64.dll"
	}
}

func SQLAllocHandle(handleType SQLSMALLINT, inputHandle SQLHANDLE, outputHandle *SQLHANDLE) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall(procSQLAllocHandle.Addr(), 3, uintptr(handleType), uintptr(inputHandle), uintptr(unsafe.Pointer(outputHandle)))
	ret = SQLRETURN(r0)
	return
}

func SQLBindCol(statementHandle SQLHSTMT, columnNumber SQLUSMALLINT, targetType SQLSMALLINT, targetValuePtr SQLPOINTER, bufferLength SQLLEN, vallen *SQLLEN) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall6(procSQLBindCol.Addr(), 6, uintptr(statementHandle), uintptr(columnNumber), uintptr(targetType), uintptr(targetValuePtr), uintptr(bufferLength), uintptr(unsafe.Pointer(vallen)))
	ret = SQLRETURN(r0)
	return
}

func SQLBindParameter(statementHandle SQLHSTMT, parameterNumber SQLUSMALLINT, inputOutputType SQLSMALLINT, valueType SQLSMALLINT, parameterType SQLSMALLINT, columnSize SQLULEN, decimalDigits SQLSMALLINT, parameterValue SQLPOINTER, bufferLength SQLLEN, ind *SQLLEN) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall12(procSQLBindParameter.Addr(), 10, uintptr(statementHandle), uintptr(parameterNumber), uintptr(inputOutputType), uintptr(valueType), uintptr(parameterType), uintptr(columnSize), uintptr(decimalDigits), uintptr(parameterValue), uintptr(bufferLength), uintptr(unsafe.Pointer(ind)), 0, 0)
	ret = SQLRETURN(r0)
	return
}

func SQLCloseCursor(statementHandle SQLHSTMT) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall(procSQLCloseCursor.Addr(), 1, uintptr(statementHandle), 0, 0)
	ret = SQLRETURN(r0)
	return
}

func SQLDescribeCol(statementHandle SQLHSTMT, columnNumber SQLUSMALLINT, columnName *SQLWCHAR, bufferLength SQLSMALLINT, nameLengthPtr *SQLSMALLINT, dataTypePtr *SQLSMALLINT, columnSizePtr *SQLULEN, decimalDigitsPtr *SQLSMALLINT, nullablePtr *SQLSMALLINT) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall9(procSQLDescribeColW.Addr(), 9, uintptr(statementHandle), uintptr(columnNumber), uintptr(unsafe.Pointer(columnName)), uintptr(bufferLength), uintptr(unsafe.Pointer(nameLengthPtr)), uintptr(unsafe.Pointer(dataTypePtr)), uintptr(unsafe.Pointer(columnSizePtr)), uintptr(unsafe.Pointer(decimalDigitsPtr)), uintptr(unsafe.Pointer(nullablePtr)))
	ret = SQLRETURN(r0)
	return
}

func SQLDescribeParam(statementHandle SQLHSTMT, parameterNumber SQLUSMALLINT, dataTypePtr *SQLSMALLINT, parameterSizePtr *SQLULEN, decimalDigitsPtr *SQLSMALLINT, nullablePtr *SQLSMALLINT) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall6(procSQLDescribeParam.Addr(), 6, uintptr(statementHandle), uintptr(parameterNumber), uintptr(unsafe.Pointer(dataTypePtr)), uintptr(unsafe.Pointer(parameterSizePtr)), uintptr(unsafe.Pointer(decimalDigitsPtr)), uintptr(unsafe.Pointer(nullablePtr)))
	ret = SQLRETURN(r0)
	return
}

func SQLDisconnect(connectionHandle SQLHDBC) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall(procSQLDisconnect.Addr(), 1, uintptr(connectionHandle), 0, 0)
	ret = SQLRETURN(r0)
	return
}

func SQLDriverConnect(connectionHandle SQLHDBC, windowHandle SQLHWND, inConnectionString *SQLWCHAR, stringLength1 SQLSMALLINT, outConnectionString *SQLWCHAR, bufferLength SQLSMALLINT, stringLength2Ptr *SQLSMALLINT, driverCompletion SQLUSMALLINT) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall9(procSQLDriverConnectW.Addr(), 8, uintptr(connectionHandle), uintptr(windowHandle), uintptr(unsafe.Pointer(inConnectionString)), uintptr(stringLength1), uintptr(unsafe.Pointer(outConnectionString)), uintptr(bufferLength), uintptr(unsafe.Pointer(stringLength2Ptr)), uintptr(driverCompletion), 0)
	ret = SQLRETURN(r0)
	return
}

func SQLEndTran(handleType SQLSMALLINT, handle SQLHANDLE, completionType SQLSMALLINT) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall(procSQLEndTran.Addr(), 3, uintptr(handleType), uintptr(handle), uintptr(completionType))
	ret = SQLRETURN(r0)
	return
}

func SQLExecute(statementHandle SQLHSTMT) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall(procSQLExecute.Addr(), 1, uintptr(statementHandle), 0, 0)
	ret = SQLRETURN(r0)
	return
}

func SQLFetch(statementHandle SQLHSTMT) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall(procSQLFetch.Addr(), 1, uintptr(statementHandle), 0, 0)
	ret = SQLRETURN(r0)
	return
}

func SQLFreeHandle(handleType SQLSMALLINT, handle SQLHANDLE) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall(procSQLFreeHandle.Addr(), 2, uintptr(handleType), uintptr(handle), 0)
	ret = SQLRETURN(r0)
	return
}

func SQLGetData(statementHandle SQLHSTMT, colOrParamNum SQLUSMALLINT, targetType SQLSMALLINT, targetValuePtr SQLPOINTER, bufferLength SQLLEN, vallen *SQLLEN) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall6(procSQLGetData.Addr(), 6, uintptr(statementHandle), uintptr(colOrParamNum), uintptr(targetType), uintptr(targetValuePtr), uintptr(bufferLength), uintptr(unsafe.Pointer(vallen)))
	ret = SQLRETURN(r0)
	return
}

func SQLGetDiagRec(handleType SQLSMALLINT, handle SQLHANDLE, recNumber SQLSMALLINT, sqlState *SQLWCHAR, nativeErrorPtr *SQLINTEGER, messageText *SQLWCHAR, bufferLength SQLSMALLINT, textLengthPtr *SQLSMALLINT) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall9(procSQLGetDiagRecW.Addr(), 8, uintptr(handleType), uintptr(handle), uintptr(recNumber), uintptr(unsafe.Pointer(sqlState)), uintptr(unsafe.Pointer(nativeErrorPtr)), uintptr(unsafe.Pointer(messageText)), uintptr(bufferLength), uintptr(unsafe.Pointer(textLengthPtr)), 0)
	ret = SQLRETURN(r0)
	return
}

func SQLNumParams(statementHandle SQLHSTMT, parameterCountPtr *SQLSMALLINT) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall(procSQLNumParams.Addr(), 2, uintptr(statementHandle), uintptr(unsafe.Pointer(parameterCountPtr)), 0)
	ret = SQLRETURN(r0)
	return
}

func SQLNumResultCols(statementHandle SQLHSTMT, columnCountPtr *SQLSMALLINT) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall(procSQLNumResultCols.Addr(), 2, uintptr(statementHandle), uintptr(unsafe.Pointer(columnCountPtr)), 0)
	ret = SQLRETURN(r0)
	return
}

func SQLPrepare(statementHandle SQLHSTMT, statementText *SQLWCHAR, textLength SQLINTEGER) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall(procSQLPrepareW.Addr(), 3, uintptr(statementHandle), uintptr(unsafe.Pointer(statementText)), uintptr(textLength))
	ret = SQLRETURN(r0)
	return
}

func SQLRowCount(statementHandle SQLHSTMT, rowCountPtr *SQLLEN) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall(procSQLRowCount.Addr(), 2, uintptr(statementHandle), uintptr(unsafe.Pointer(rowCountPtr)), 0)
	ret = SQLRETURN(r0)
	return
}

func SQLSetEnvAttr(environmentHandle SQLHENV, attribute SQLINTEGER, valuePtr SQLPOINTER, stringLength SQLINTEGER) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall6(procSQLSetEnvAttr.Addr(), 4, uintptr(environmentHandle), uintptr(attribute), uintptr(valuePtr), uintptr(stringLength), 0, 0)
	ret = SQLRETURN(r0)
	return
}

func SQLSetConnectAttr(connectionHandle SQLHDBC, attribute SQLINTEGER, valuePtr SQLPOINTER, stringLength SQLINTEGER) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall6(procSQLSetConnectAttrW.Addr(), 4, uintptr(connectionHandle), uintptr(attribute), uintptr(valuePtr), uintptr(stringLength), 0, 0)
	ret = SQLRETURN(r0)
	return
}

func SQLColAttribute(statementHandle SQLHSTMT, ColumnNumber SQLUSMALLINT, FieldIdentifier SQLUSMALLINT, CharacterAttributePtr SQLPOINTER, BufferLength SQLSMALLINT, StringLengthPtr *SQLSMALLINT, NumericAttributePtr SQLPOINTER) (ret SQLRETURN) {
	r0, _, _ := syscall.Syscall9(procSQLColAttribute.Addr(), 7, uintptr(statementHandle), uintptr(ColumnNumber), uintptr(FieldIdentifier), uintptr(CharacterAttributePtr), uintptr(BufferLength), uintptr(unsafe.Pointer(StringLengthPtr)), uintptr(NumericAttributePtr), 0, 0)
	ret = SQLRETURN(r0)
	return
}
