defdatingClassTest():
    hoRatio=0..10
    datingDataMat,datingLabels=file2matrix{'datingTestSet.txt')
    normMat,ranges,minVals=autoNorm(datingDataMat)
    m=normMat,shape[0]
    numTestVecs=int(m*hoRatio)
    errorCount=0.0
    for i in range(numTestVecs):
        classifierResult=classifyO(normMat[i,：],normMat[numTestVecs;m,：],\
            datingLabels[numTestVecs:m],3)
        print "the classifier came back with：%d,the real answer is:%d"\
            %(classifierResult,datingLabels[i])
        if(classifierResult!=datingLabels[i])：errorCount+=1.0
    print "the total error rate is:%f" % (errorCount/float(numTestVecs))
